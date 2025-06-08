package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kpl-base/application/service"
	"kpl-base/command"
	domain_user "kpl-base/domain/user"
	"kpl-base/infrastructure/adapter/file_storage"
	"kpl-base/infrastructure/database/config"
	infrastructure_refresh_token "kpl-base/infrastructure/database/refresh_token"
	"kpl-base/infrastructure/database/transaction"
	infrastructure_user "kpl-base/infrastructure/database/user"
	"kpl-base/presentation/controller"
	"kpl-base/presentation/middleware"
	"kpl-base/presentation/route"
	"log"
	"os"
)

func args(db *gorm.DB) bool {
	if len(os.Args) > 1 {
		flag := command.Commands(db)
		return flag
	}

	return true
}

func run(server *gin.Engine) {
	server.Static("/assets", "./assets")

	if os.Getenv("IS_LOGGER") == "true" {
		route.LoggerRoute(server)
	}

	port := os.Getenv("GOLANG_PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "0.0.0.0:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}

func main() {
	db := config.SetUpDatabaseConnection()

	jwtService := service.NewJWTService()

	transactionRepository := transaction.NewRepository(db)
	userRepository := infrastructure_user.NewRepository(transactionRepository)
	refreshTokenRepository := infrastructure_refresh_token.NewRepository(transactionRepository)

	fileStorage := file_storage.NewLocalAdapter()

	userDomainService := domain_user.NewService(fileStorage)

	userService := service.NewUserService(userRepository, refreshTokenRepository, *userDomainService, jwtService, transactionRepository)

	userController := controller.NewUserController(userService)

	defer config.CloseDatabaseConnection(db)

	if !args(db) {
		return
	}

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	route.UserRoute(server, userController, jwtService)

	run(server)
}
