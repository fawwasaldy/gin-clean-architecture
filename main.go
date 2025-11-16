package main

import (
	"log"
	"os"

	"github.com/fawwasaldy/gin-clean-architecture/command"
	"github.com/fawwasaldy/gin-clean-architecture/internal/infrastructure/database/config"
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/middleware"
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/route"
	"github.com/fawwasaldy/gin-clean-architecture/platform/provider"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
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
	var (
		injector = do.New()
	)

	provider.RegisterDependencies(injector)

	db := do.MustInvoke[*gorm.DB](injector)

	defer config.CloseDatabaseConnection(db)

	if !args(db) {
		return
	}

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	do.ProvideValue(injector, server)

	route.RegisterRoutes(injector)

	run(server)
}
