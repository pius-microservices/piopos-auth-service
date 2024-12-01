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
	Token string `json:"token"`
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

	token, err := jwt.CreateToken()

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"data": user, "accessToken": tokenResponse{Token: token}}, 200
}
