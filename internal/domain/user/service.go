package user

import (
	"fmt"
	"mime/multipart"

	"github.com/fawwasaldy/gin-clean-architecture/internal/domain/port"
	"github.com/samber/do/v2"

	"github.com/google/uuid"
)

type Service struct {
	fileStorage port.FileStoragePort
}

func NewService(injector do.Injector) *Service {
	fileStorage := do.MustInvoke[port.FileStoragePort](injector)
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
