package Middleware

import (
	"io"
	"log"
	"mime"
	"net/http"

	context "github.com/ines-mgg/LetsGoBack/Context"
)

// UploadValidatorMiddleware validates uploaded files based on the provided options.
// It checks the MIME type, file size, and whether multiple files are allowed.
// If validation fails, it responds with a bad request error.
// If validation succeeds, it stores the uploaded file(s) in the context for further processing.
// The options include:
// - MaxFileSize: Maximum allowed file size in bytes.
// - AllowedMIMEs: List of allowed MIME types.
// - Field: The form field name for the uploaded file(s).
// - Multiple: Whether to allow multiple files to be uploaded.
// - MaxMemory: Maximum memory in bytes to use for file uploads.
func UploadValidatorMiddleware(opts UploadValidationOptions) Middleware {
	return func(next context.HandlerFunc) context.HandlerFunc {
		return func(c *context.Context) {
			var files []*context.UploadedFile
			var err error

			if opts.Multiple {
				files, err = c.GetUploadedFiles(opts.Field, opts.MaxMemory)
			} else {
				var f *context.UploadedFile
				f, err = c.GetUploadedFile(opts.Field, opts.MaxMemory)
				if f != nil {
					files = []*context.UploadedFile{f}
				}
			}

			if err != nil {
				log.Printf("✖️ Error getting uploaded file(s): %v", err)
				c.ErrorBadRequest("upload error")
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
					c.ErrorBadRequest("invalid content type")
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
					c.ErrorBadRequest("invalid mime type")
					return
				}

				if opts.MaxFileSize > 0 && file.Size > opts.MaxFileSize {
					log.Printf("✖️ File too large: %d bytes (max: %d)", file.Size, opts.MaxFileSize)
					c.ErrorBadRequest("file too large")
					return
				}
			}

			// Store the files in the context for further use
			if opts.Multiple {
				c.Set("uploadedFiles", files)
			} else {
				c.Set("uploadedFile", files[0])
			}
			log.Printf("✔️ Uploaded file(s) validated successfully: %v", files)
			next(c)
		}
	}
}
