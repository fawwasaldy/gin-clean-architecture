package user

import (
	"github.com/fawwasaldy/gin-clean-architecture/internal/application/service"
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/controller"
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/middleware"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

func Route(injector do.Injector) {
	baseRoute := do.MustInvoke[*gin.RouterGroup](injector)
	jwtService := do.MustInvoke[service.JWTService](injector)
	userController := do.MustInvoke[controller.UserController](injector)

	userGroup := baseRoute.Group("/user")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		userGroup.POST("/refresh-token", userController.RefreshToken)
		userGroup.POST("/logout", middleware.Authenticate(jwtService), userController.Logout)
		userGroup.GET("/", middleware.Authenticate(jwtService), userController.GetAll)
		userGroup.PATCH("/", middleware.Authenticate(jwtService), userController.Update)
		userGroup.DELETE("/", middleware.Authenticate(jwtService), userController.Delete)
	}
}
