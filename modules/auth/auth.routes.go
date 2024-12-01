package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.Engine, controller *authController, prefix string) {
	authGroup := router.Group(prefix + "/auth")
	{
		authGroup.POST("/signin", func(ctx *gin.Context) {
			controller.SignIn(ctx)
		})
	}
}
