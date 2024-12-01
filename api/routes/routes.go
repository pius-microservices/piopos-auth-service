package routes

import (
	_ "github.com/pius-microservices/piopos-auth-service/docs"
	"github.com/pius-microservices/piopos-auth-service/modules/auth"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	APIPrefix = "/api/auth-service"
)

func homeHandler(ctx *gin.Context) {
	ctx.JSON(404, gin.H{
		"status":  404,
		"message": "Sorry! Page not found",
	})
}

func RouteApp(router *gin.Engine) error {
	router.GET(APIPrefix+"/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET(APIPrefix, homeHandler)

	auth.AuthRoutesModule(router, APIPrefix)

	return nil
}
