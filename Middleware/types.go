package Middleware

import context "lets-go-back/Context"

type Middleware func(context.HandlerFunc) context.HandlerFunc

type UploadValidationOptions struct {
	MaxFileSize  int64
	AllowedMIMEs []string
	Field        string
	Multiple     bool
}
