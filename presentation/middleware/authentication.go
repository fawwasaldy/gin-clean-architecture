package middleware

import (
	"github.com/gin-gonic/gin"
	"kpl-base/application/service"
	"kpl-base/presentation"
	"kpl-base/presentation/message"
	"net/http"
	"strings"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response := presentation.BuildResponseFailed(message.FailedProcessRequest, message.FailedTokenNotFound, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !strings.Contains(authHeader, "Bearer ") {
			response := presentation.BuildResponseFailed(message.FailedProcessRequest, message.FailedTokenNotValid, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := presentation.BuildResponseFailed(message.FailedProcessRequest, message.FailedTokenNotValid, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := presentation.BuildResponseFailed(message.FailedProcessRequest, message.FailedDeniedAccess, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := presentation.BuildResponseFailed(message.FailedProcessRequest, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("token", authHeader)
		ctx.Set("user_id", userId)
		ctx.Next()
	}
}
