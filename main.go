package main

import (
	"log"
	"os"

	"github.com/fawwasaldy/gin-clean-architecture/command"
	"github.com/fawwasaldy/gin-clean-architecture/internal/application/service"
	"github.com/fawwasaldy/gin-clean-architecture/internal/domain/user"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/adapter/file_storage"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/config"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/repository"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/transaction"
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/controller"
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/middleware"
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/route"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	userRepository := repository.NewUserRepository(transactionRepository)
	refreshTokenRepository := repository.NewRefreshTokenRepository(transactionRepository)

	fileStorage := file_storage.NewLocalAdapter()

	userDomainService := user.NewService(fileStorage)

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
