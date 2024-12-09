package auth

import (
	"github.com/pius-microservices/piopos-auth-service/interfaces"
	"github.com/pius-microservices/piopos-auth-service/middlewares"
	"github.com/pius-microservices/piopos-auth-service/package/utils"

	"github.com/gin-gonic/gin"
)

type authService struct {
	repo interfaces.AuthRepo
}

func NewService(repo interfaces.AuthRepo) *authService {
	return &authService{repo}
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (service *authService) SignIn(email string, password string) (gin.H, int) {
	user, err := service.repo.FetchUserByEmail(email)

	if err != nil {
		return gin.H{"status": 401, "message": "Email or password is incorrect"}, 401
	}

	if !utils.CheckPassword(password, user.Password) {
		return gin.H{"status": 401, "message": "Email or password is incorrect"}, 401
	}

	jwt := middlewares.NewToken(user.ID)

	accessToken, err := jwt.CreateToken()

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	refreshToken, err := service.repo.CreateRefreshToken(user.ID)

	if err != nil {
		return gin.H{"status": 500, "message": "Failed to generate refresh token"}, 500
	}

	return gin.H{"data": user, "token": tokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}}, 200
}