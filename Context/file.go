package Context

import (
	"io"
	"os"
	"path/filepath"
)

// setDefaultMaxMemory sets a default maximum memory limit for file uploads.
// If maxMemory is 0, it defaults to 32 MB (32 << 20 bytes).
// This function is used to ensure that file uploads do not exceed a reasonable size,
// preventing excessive memory usage during file handling.
// It returns the maximum memory limit to be used for file uploads.
func setDefaultMaxMemory(maxMemory int64) int64 {
	if maxMemory == 0 {
		return 32 << 20 // 32 MB
	}
	return maxMemory
}

// GetUploadedFile retrieves a single uploaded file from the request form.
// It uses the specified field name to find the file in the form data.
// If the file is found, it returns an UploadedFile struct containing the file,
// its header, filename, and size. If the file is not found or there is an error,
// it returns an error.
// The maxMemory parameter specifies the maximum memory to use when parsing the form data.
// If maxMemory is 0, it defaults to 32 MB.
func (c *Context) GetUploadedFile(field string, maxMemory int64) (*UploadedFile, error) {
	maxMemory = setDefaultMaxMemory(maxMemory)
	file, fileHeader, err := c.Request.FormFile(field)
	if err != nil {
		return nil, err
	}
	err = c.Request.ParseMultipartForm(maxMemory)

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

// GetUploadedFiles retrieves multiple uploaded files from the request form.
// It uses the specified field name to find the files in the form data.
// If files are found, it returns a slice of UploadedFile structs, each containing
// the file, its header, filename, and size. If there is an error during retrieval,
// it returns an error.
// The maxMemory parameter specifies the maximum memory to use when parsing the form data.
// If maxMemory is 0, it defaults to 32 MB.
func (c *Context) GetUploadedFiles(field string, maxMemory int64) ([]*UploadedFile, error) {
	maxMemory = setDefaultMaxMemory(maxMemory)
	err := c.Request.ParseMultipartForm(maxMemory)
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

// SaveFile saves an uploaded file to the specified path.
// It creates the necessary directories if they do not exist.
// The function takes an UploadedFile struct and a path as parameters.
// It returns an error if there is an issue creating the directories or writing the file.
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

// DeleteFile deletes a file at the specified path.
// If the file does not exist, it returns nil.
// If the file exists, it attempts to remove it and returns any error encountered.
func (c *Context) DeleteFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(path)
}
