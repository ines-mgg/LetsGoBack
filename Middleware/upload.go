package Middleware

import (
	"fmt"
	"io"
	context "lets-go-back/Context"
	"log"
	"mime"
	"net/http"
	"strings"
)

func UploadLoggerMiddleware(next context.HandlerFunc) context.HandlerFunc {
	return func(c *context.Context) {
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
	return func(next context.HandlerFunc) context.HandlerFunc {
		return func(c *context.Context) {
			var files []*context.UploadedFile
			var err error

			if opts.Multiple {
				files, err = c.GetUploadedFiles(opts.Field)
			} else {
				var f *context.UploadedFile
				f, err = c.GetUploadedFile(opts.Field)
				if f != nil {
					files = []*context.UploadedFile{f}
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
