package provider

import (
	"github.com/fawwasaldy/gin-clean-architecture/internal/application/service"
	"github.com/fawwasaldy/gin-clean-architecture/internal/domain/refresh_token"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/config"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/repository"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/transaction"
	"github.com/fawwasaldy/gin-clean-architecture/platform/provider/user"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func RegisterDependencies(injector do.Injector) {
	InitDatabase(injector)
	InitJWTService(injector)
	InitRefreshTokenRepository(injector)
	InitTransactionRepository(injector)

	RegisterAdapterDependencies(injector)
	user.RegisterDependencies(injector)
}

func InitDatabase(injector do.Injector) {
	do.Provide(injector, func(injector do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func InitJWTService(injector do.Injector) {
	do.Provide(injector, func(injector do.Injector) (service.JWTService, error) {
		return service.NewJWTService(), nil
	})
}

func InitRefreshTokenRepository(injector do.Injector) {
	do.Provide(injector, func(injector do.Injector) (refresh_token.Repository, error) {
		return repository.NewRefreshTokenRepository(injector), nil
	})
}

func InitTransactionRepository(injector do.Injector) {
	do.Provide(injector, func(injector do.Injector) (*transaction.Repository, error) {
		return transaction.NewRepository(injector), nil
	})
}
