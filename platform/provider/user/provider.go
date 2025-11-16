package user

import (
	"github.com/fawwasaldy/gin-clean-architecture/internal/application/service"
	"github.com/fawwasaldy/gin-clean-architecture/internal/domain/user"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/repository"
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/controller"
	"github.com/samber/do/v2"
)

func RegisterDependencies(injector do.Injector) {
	do.Provide(injector, func(injector do.Injector) (user.Repository, error) {
		return repository.NewUserRepository(injector), nil
	})
	do.Provide(injector, func(injector do.Injector) (*user.Service, error) {
		return user.NewService(injector), nil
	})
	do.Provide(injector, func(injector do.Injector) (service.UserService, error) {
		return service.NewUserService(injector), nil
	})
	do.Provide(injector, func(injector do.Injector) (controller.UserController, error) {
		return controller.NewUserController(injector), nil
	})
}
