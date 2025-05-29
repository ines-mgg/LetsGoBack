package Context

import (
	"mime/multipart"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request

	Path   string
	Method string
	Status int

	Params map[string]string
	Data   map[string]any
}

type UploadedFile struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	Filename   string
	Size       int64
}

type HandlerFunc func(*Context)
