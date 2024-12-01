package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pius-microservices/piopos-auth-service/config"
	"github.com/pius-microservices/piopos-auth-service/interfaces"
	"github.com/pius-microservices/piopos-auth-service/package/database/models"
)

type authRepo struct{}

func NewRepo() interfaces.AuthRepo {
	return &authRepo{}
}

func (repo *authRepo) FetchUserByEmail(email string) (*models.User, error) {
	envCfg := config.LoadConfig()
	url := fmt.Sprintf(envCfg.UserServiceBaseURL + envCfg.GetUserByEmail, email)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var userResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}

	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		return nil, err
	}

	if userResponse.Status != 200 {
		return nil, fmt.Errorf(userResponse.Message)
	}

	return &userResponse.Data, nil
}
