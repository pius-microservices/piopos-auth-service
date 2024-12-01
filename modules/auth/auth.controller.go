package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pius-microservices/piopos-auth-service/interfaces"
	"github.com/pius-microservices/piopos-auth-service/package/database/models"
)

type authController struct {
	service interfaces.AuthService
}

func NewController(service interfaces.AuthService) *authController {
	return &authController{service}
}

// SignIn godoc
// @Summary Login as an authenticated user
// @Description Login with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userData body models.SignInRequest true "User data"
// @Success 200
// @Failure 401
// @Failure 500
// @Router /api/auth-service/auth/signin [post]
func (controller *authController) SignIn(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var userData models.SignInRequest

	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	if userData.Password == "" {
        ctx.JSON(401, gin.H{"error": "You forget to enter your password"})
        return
    }

	responseData, status := controller.service.SignIn(userData.Email, userData.Password)

	ctx.JSON(status, responseData)
}
