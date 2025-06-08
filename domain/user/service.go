package user

import (
	"fmt"
	"github.com/google/uuid"
	"kpl-base/domain/port"
	"mime/multipart"
)

type Service struct {
	fileStorage port.FileStoragePort
}

func NewService(fileStorage port.FileStoragePort) *Service {
	return &Service{
		fileStorage: fileStorage,
	}
}

func (s *Service) UploadImage(image *multipart.FileHeader) (filename string, err error) {
	imageId := uuid.New()
	ext := s.fileStorage.GetExtension(image.Filename)

	filename = fmt.Sprintf("profile/%s.%s", imageId.String(), ext)
	if err = s.fileStorage.UploadFile(image, filename); err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	return filename, nil
}
