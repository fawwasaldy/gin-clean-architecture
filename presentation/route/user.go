package route

import (
	"github.com/gin-gonic/gin"
	"kpl-base/application/service"
	"kpl-base/presentation/controller"
	"kpl-base/presentation/middleware"
)

func UserRoute(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	userGroup := route.Group("/api/user")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.GET("/me", middleware.Authenticate(jwtService), userController.Me)
		userGroup.POST("/refresh-token", userController.RefreshToken)
		userGroup.GET("/", middleware.Authenticate(jwtService), userController.GetAll)
		userGroup.PATCH("/", middleware.Authenticate(jwtService), userController.Update)
		userGroup.DELETE("/", middleware.Authenticate(jwtService), userController.Delete)
	}
}
