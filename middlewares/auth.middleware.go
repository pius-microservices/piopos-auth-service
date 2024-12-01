package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authorization header required", "error": true})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "invalid header type", "error": true})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		checkToken, err := VerifyToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "error": true})
			ctx.Abort()
			return
		}

		ctx.Set("id", checkToken.UserId)

		ctx.Next()
	}
}
