// Package core provides the implementation of a lightweight HTTP router
// with support for dynamic routes, middleware, and route grouping.
// This file, middleware.go, defines the Middleware type and common middleware functions,
// such as logging and panic recovery, to enhance the request handling process.
package core

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Types for middleware.go
type Middleware func(HandlerFunc) HandlerFunc

type UploadValidationOptions struct {
	MaxFileSize  int64
	AllowedMIMEs []string
	Field        string
	Multiple     bool
}

func logLevel(status int) string {
	switch {
	case status >= 500:
		return "[ERROR]"
	case status >= 400:
		return "[WARN]"
	default:
		return "[INFO]"
	}
}

func LoggerMiddleware(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		start := time.Now()

		next(c)

		duration := time.Since(start)
		status := c.Status()
		method := c.method
		path := c.path

		reqID, _ := c.Get("request_id")
		reqIDStr, _ := reqID.(string)

		if status == 0 {
			status = 200 // fallback
		}

		if reqID == "" {
			reqID = "n/a"
		}

		log.Printf("%s [%s] [%s] %s - %d (%s) - reqID: %s",
			logLevel(status),
			time.Now().Format("2006-01-02 15:04:05"),
			method,
			path,
			status,
			duration.Truncate(time.Millisecond),
			reqIDStr,
		)
	}
}

func applyMiddleware(h HandlerFunc, middleware []Middleware) HandlerFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}

func RecoverMiddleware(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				errorID := generateErrorID()
				log.Printf("[PANIC][%s] %s %s - %v\n%s", errorID, c.Request.Method, c.Request.URL.Path, err, string(debug.Stack()))
				c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("Internal server error. Error ID: %s", errorID))
			}
		}()
		next(c)
	}
}

func SessionAuthMiddleware(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		userID, ok := GetUserID(c.Request)
		if !ok {
			c.JSON(401, map[string]string{"error": "Unauthorized"})
			return
		}

		c.Set("userID", userID)
		next(c)
	}
}

func JWTAuthMiddleware(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		cookie, err := c.Request.Cookie("auth_token")
		if err != nil {
			c.JSON(401, map[string]string{"error": "missing token"})
			return
		}

		userID, err := ParseJWT(cookie.Value)
		if err != nil {
			c.JSON(401, map[string]string{"error": "invalid token"})
			return
		}

		c.Set("userID", userID)
		next(c)
	}
}

func CORSMiddleware(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		origin := c.Request.Header.Get("Origin")

		if origin == "http://localhost:5500" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Vary", "Origin")
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(http.StatusOK)
			return
		}

		next(c)
	}
}

func UploadLoggerMiddleware(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		if c.Request.Method == http.MethodPost &&
			strings.HasPrefix(c.Request.Header.Get("Content-Type"), "multipart/form-data") {

			err := c.Request.ParseMultipartForm(32 << 20) // 32 MB
			if err == nil && c.Request.MultipartForm != nil {
				for field, files := range c.Request.MultipartForm.File {
					for _, fh := range files {
						log.Printf("Upload - field: %s, filename: %s, size: %d bytes",
							field, fh.Filename, fh.Size)
					}
				}
			} else {
				log.Printf("Upload - failed to parse multipart form: %v", err)
			}
		}

		next(c)
	}
}

func UploadValidatorMiddleware(opts UploadValidationOptions) Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			var files []*UploadedFile
			var err error

			if opts.Multiple {
				files, err = c.GetUploadedFiles(opts.Field)
			} else {
				var f *UploadedFile
				f, err = c.GetUploadedFile(opts.Field)
				if f != nil {
					files = []*UploadedFile{f}
				}
			}

			if err != nil {
				log.Printf("✖️ Error getting uploaded file(s): %v", err)
				c.JSON(http.StatusBadRequest, map[string]string{"error": "upload error"})
				return
			}

			for _, file := range files {
				buf := make([]byte, 512)
				n, _ := file.File.Read(buf)
				file.File.Seek(0, io.SeekStart)

				var contentType string

				if file.FileHeader != nil {
					contentType = file.FileHeader.Header.Get("Content-Type")
				}

				if contentType == "" {
					contentType = http.DetectContentType(buf[:n])
				}

				mimeType, _, err := mime.ParseMediaType(contentType)
				if err != nil {
					log.Printf("✖️ Failed to parse MIME type: %v", err)
					c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid content type"})
					return
				}

				allowed := false
				for _, allowedType := range opts.AllowedMIMEs {
					if mimeType == allowedType {
						allowed = true
						break
					}
				}

				if !allowed {
					log.Printf("✖️ MIME type not allowed: %s (allowed: %v)", mimeType, opts.AllowedMIMEs)
					c.JSON(http.StatusBadRequest, map[string]string{
						"error": fmt.Sprintf("invalid mime type: %s", mimeType),
					})
					return
				}

				if opts.MaxFileSize > 0 && file.Size > opts.MaxFileSize {
					log.Printf("✖️ File too large: %d bytes (max: %d)", file.Size, opts.MaxFileSize)
					c.JSON(http.StatusBadRequest, map[string]string{
						"error": fmt.Sprintf("file too large: %d bytes", file.Size),
					})
					return
				}
			}

			next(c)
		}
	}
}

func RequestIDMiddleware(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Set("request_id", requestID)

		next(c)
	}
}

func ErrorRecoveryMiddleware() Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Printf("[ERROR] [%s] Panic recovered: %v - path: %s", c.RequestID(), rec, c.Path())

					c.JSON(http.StatusInternalServerError, map[string]string{
						"error": "internal server error",
					})
				}
			}()

			next(c)
		}
	}
}
