package route

import (
	"github.com/gin-gonic/gin"
	"kpl-base/application/service"
	"kpl-base/presentation/controller"
)

func UserRoute(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	userGroup := route.Group("/users")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.GET("/me", userController.Me)
		userGroup.POST("/refresh-token", userController.RefreshToken)
		userGroup.PUT("/", userController.Update)
		userGroup.DELETE("/", userController.Delete)
	}
}
