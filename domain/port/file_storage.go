package port

import "mime/multipart"

type (
	FileStoragePort interface {
		UploadFile(file *multipart.FileHeader, path string) error
		GetExtension(filename string) string
	}
)
