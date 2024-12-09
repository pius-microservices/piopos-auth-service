package interfaces

import (
	"github.com/pius-microservices/piopos-auth-service/package/database/models"

	"github.com/gin-gonic/gin"
)

type AuthRepo interface {
	FetchUserByEmail(email string) (*models.User, error)
	CreateRefreshToken(userId string) (string, error)
}

type AuthService interface {
	SignIn(email string, password string) (gin.H, int)
}