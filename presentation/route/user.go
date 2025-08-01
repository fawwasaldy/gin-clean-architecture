package route

import (
	"gin-clean-architecture/application/service"
	"gin-clean-architecture/presentation/controller"
	"gin-clean-architecture/presentation/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	userGroup := route.Group("/api/user")
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
