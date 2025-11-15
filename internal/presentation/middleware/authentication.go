package middleware

import (
	"net/http"
	"strings"

	"github.com/fawwasaldy/gin-clean-architecture/internal/application/service"
	"github.com/fawwasaldy/gin-clean-architecture/internal/presentation/message"
	"github.com/fawwasaldy/gin-clean-architecture/platform/response"
	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			res := response.BuildResponseFailed(message.FailedProcessRequest, message.FailedTokenNotFound, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			res := response.BuildResponseFailed(message.FailedProcessRequest, message.FailedTokenNotValid, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			res := response.BuildResponseFailed(message.FailedProcessRequest, message.FailedTokenNotValid, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		if !token.Valid {
			res := response.BuildResponseFailed(message.FailedProcessRequest, message.FailedDeniedAccess, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		userId, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			res := response.BuildResponseFailed(message.FailedProcessRequest, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		ctx.Set("token", authHeader)
		ctx.Set("user_id", userId)
		ctx.Next()
	}
}
