package Context

import (
	"mime/multipart"
	"net/http"
)

// Context represents the context of an HTTP request.
// It contains the request and response writer, as well as various properties
// such as the request path, method, status code, parameters, and data.
// This context is used to pass information between middleware and handlers in the application.
// It allows handlers to access request data, set response headers, and manage the response status.
// The Context struct is designed to be lightweight and efficient, providing a simple interface
// for handling HTTP requests in a web application.
// It is typically created at the beginning of request processing and passed through the middleware chain
// and to the final handler.
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request

	Path   string
	Method string
	Status int

	Params map[string]string
	Data   map[string]any
}

// UploadedFile represents a file that has been uploaded in an HTTP request.
// It contains the file itself, its header, filename, and size.
// This struct is used to handle file uploads in web applications.
// It provides a convenient way to access the uploaded file's content and metadata.
// The File field is of type multipart.File, which allows reading the file's content.
// The FileHeader field contains metadata about the uploaded file, such as its name and size.
// The Filename field is a string representation of the file's name, and Size is the size of the file in bytes.
// This struct is typically used in conjunction with multipart form data handling in HTTP requests.
// It is useful for applications that allow users to upload files, such as images, documents, or other types of data.
type UploadedFile struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	Filename   string
	Size       int64
}

// HandlerFunc is a function type that defines the signature for HTTP handlers.
// It takes a pointer to a Context as an argument, allowing access to the request and response data.
type HandlerFunc func(*Context)
