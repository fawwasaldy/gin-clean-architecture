package file_storage

import (
	"fmt"
	"io"
	"kpl-base/domain/port"
	"mime/multipart"
	"os"
	"strings"
)

const Path = "assets"

type localAdapter struct{}

func NewLocalAdapter() port.FileStoragePort {
	return &localAdapter{}
}

func (l localAdapter) UploadFile(file *multipart.FileHeader, path string) error {
	parts := strings.Split(path, "/")
	fileID := parts[1]
	dirPath := fmt.Sprintf("%s/%s", Path, fileID)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
	}

	filePath := fmt.Sprintf("%s/%s", dirPath, fileID)

	uploadedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadedFile.Close()

	targetFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer targetFile.Close()

	_, err = io.Copy(targetFile, uploadedFile)
	if err != nil {
		return err
	}

	return nil
}

func (l localAdapter) GetExtension(filename string) string {
	return strings.Split(filename, ".")[len(strings.Split(filename, "."))-1]
}
