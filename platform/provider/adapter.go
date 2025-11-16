package provider

import (
	"github.com/fawwasaldy/gin-clean-architecture/internal/domain/port"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/adapter/file_storage"
	"github.com/samber/do/v2"
)

func RegisterAdapterDependencies(injector do.Injector) {
	do.Provide(injector, func(injector do.Injector) (port.FileStoragePort, error) {
		return file_storage.NewLocalAdapter(), nil
	})
}
