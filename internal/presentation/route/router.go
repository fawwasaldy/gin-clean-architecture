package route

import (
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/route/user"
	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

func RegisterBaseRoute(injector do.Injector) {
	route := do.MustInvoke[*gin.Engine](injector)

	apiGroup := route.Group("/api")

	do.ProvideValue(injector, apiGroup)
}

func RegisterRoutes(injector do.Injector) {
	RegisterBaseRoute(injector)
	user.Route(injector)
}
