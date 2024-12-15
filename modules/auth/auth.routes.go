package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.Engine, controller *authController, prefix string) {
	authGroup := router.Group(prefix + "/auth")
	{
		authGroup.POST("/signin", func(ctx *gin.Context) {
			controller.SignIn(ctx)
		})

		authGroup.POST("/signout", func(ctx *gin.Context) {
			controller.SignOut(ctx)
		})

		authGroup.POST("/refresh-token", func(ctx *gin.Context) {
			controller.CreateNewAccessToken(ctx)
		})
	}
}
