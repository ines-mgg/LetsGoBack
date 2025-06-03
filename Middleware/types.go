package Middleware

import context "github.com/ines-mgg/LetsGoBack/Context"

// Middleware is a function that takes a context.HandlerFunc and returns a context.HandlerFunc.
// It is used to wrap handlers with additional functionality, such as authentication, logging, etc.
// The returned HandlerFunc will be executed in the context of the request.
// This allows for chaining multiple middlewares together.
// Each middleware can modify the request, response, or context before passing control to the next handler.
// The middleware can also terminate the request by writing a response directly to the context.
// This is useful for implementing cross-cutting concerns in a web application, such as error handling, request logging, or response formatting.
type Middleware func(context.HandlerFunc) context.HandlerFunc

// UploadValidationOptions defines the options for the UploadValidatorMiddleware.
// It includes the maximum file size, allowed MIME types, the field name for the uploaded file(s),
// whether multiple files are allowed, and the maximum memory to use for file uploads.
// These options allow for flexible configuration of file upload validation,
// ensuring that uploaded files meet specific criteria before being processed by the application.
// This struct is used to pass configuration options to the UploadValidatorMiddleware.
type UploadValidationOptions struct {
	MaxFileSize  int64
	AllowedMIMEs []string
	Field        string
	Multiple     bool
	MaxMemory    int64
}
