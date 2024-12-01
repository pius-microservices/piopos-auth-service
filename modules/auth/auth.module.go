package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutesModule(router *gin.Engine, prefix string) {
	authRepo := NewRepo()
	authService := NewService(authRepo)
	authController := NewController(authService)
	
	AuthRoutes(router, authController, prefix)
}