// Package core provides the implementation of a lightweight HTTP router
// with support for dynamic routes, middleware, and route grouping.
// This file, context.go, defines the Context struct, which encapsulates
// the HTTP request and response, along with helper methods for handling
// JSON responses, parameters, and error management.
package core

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Types for context.go
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request

	path   string
	method string
	status int

	Params map[string]string
	data   map[string]any
}

type UploadedFile struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	Filename   string
	Size       int64
}

// Basics functions
func generateErrorID() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Context functions
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Params:  make(map[string]string),
		path:    r.URL.Path,
		method:  r.Method,
		data:    make(map[string]any),
	}
}

func (c *Context) Get(key string) (any, bool) {
	val, ok := c.data[key]
	return val, ok
}

func (c *Context) Set(key string, value any) {
	c.data[key] = value
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) JSON(status int, data any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	if c.Status() == 0 {
		c.SetStatus(status)
		c.Writer.WriteHeader(status)
	}

	json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) BindJSON(obj interface{}) error {
	decoder := json.NewDecoder(c.Request.Body)
	return decoder.Decode(obj)
}

func (c *Context) AbortWithError(status int, err error) {
	errorID := generateErrorID()
	log.Printf("[ERROR][%s] %s %s - %v", errorID, c.Request.Method, c.Request.URL.Path, err)

	c.AbortWithStatusJSON(status, fmt.Sprintf("Something went wrong. Error ID: %s", errorID))
}

func (c *Context) AbortWithStatusJSON(status int, message string) {
	c.JSON(status, map[string]string{
		"error": message,
	})
}

func (c *Context) ErrorBadRequest(msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, msg)
}

func (c *Context) FormFile(name string) (multipart.File, *multipart.FileHeader, error) {
	return c.Request.FormFile(name)
}

func (c *Context) GetUploadedFile(field string) (*UploadedFile, error) {
	file, fileHeader, err := c.Request.FormFile(field)
	if err != nil {
		return nil, err
	}

	return &UploadedFile{
		File:       file,
		FileHeader: fileHeader,
		Filename:   fileHeader.Filename,
		Size:       fileHeader.Size,
	}, nil
}

func (c *Context) SaveFile(file *UploadedFile, path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file.File)
	return err
}

func (c *Context) GetUploadedFiles(field string) ([]*UploadedFile, error) {
	err := c.Request.ParseMultipartForm(32 << 20) // 32 MB max
	if err != nil {
		return nil, err
	}

	files := c.Request.MultipartForm.File[field]
	var uploadedFiles []*UploadedFile

	for _, f := range files {
		file, err := f.Open()
		if err != nil {
			return nil, err
		}

		uploadedFiles = append(uploadedFiles, &UploadedFile{
			Filename:   f.Filename,
			Size:       f.Size,
			FileHeader: f,
			File:       file,
		})
	}

	return uploadedFiles, nil
}

func (c *Context) RequestID() string {
	val, ok := c.Get("request_id")
	if !ok {
		return ""
	}
	if id, ok := val.(string); ok {
		return id
	}
	return ""
}

func (c *Context) Status() int {
	return c.status
}

func (c *Context) SetStatus(status int) {
	c.data["status"] = status
}

func (c *Context) Method() string {
	return c.method
}

func (c *Context) Path() string {
	return c.path
}
