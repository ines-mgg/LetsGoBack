package Context

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

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
