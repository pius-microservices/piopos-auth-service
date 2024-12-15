package auth

import (
	"bytes"
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
	url := fmt.Sprintf(envCfg.UserServiceBaseURL+envCfg.GetUserByEmail, email)

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

func (repo *authRepo) CreateRefreshToken(userId string) (string, error) {
	envCfg := config.LoadConfig()
	url := fmt.Sprintf(envCfg.UserServiceBaseURL + envCfg.CreateRefreshToken)

	requestBody := map[string]string{
		"user_id": userId,
	}

	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 201 {
		body, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("failed to generate refresh token: %s", string(body))
	}

	var tokenResponse struct {
		RefreshToken string `json:"refresh_token"`
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", nil
	}

	return tokenResponse.RefreshToken, nil
}

func (repo *authRepo) DeleteRefreshToken(userId string, refreshToken string) (string, error) {
	envCfg := config.LoadConfig()
	url := fmt.Sprintf(envCfg.UserServiceBaseURL + envCfg.DeleteRefreshToken)

	requestBody := map[string]string{
		"user_id":       userId,
		"refresh_token": refreshToken,
	}

	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("failed to delete refresh token: %s", string(body))
	}

	return "Refresh token deleted successfully", nil
}

func (repo *authRepo) ValidateRefreshToken(userId, refreshToken string) (string, error) {
	envCfg := config.LoadConfig()
	url := fmt.Sprintf(envCfg.UserServiceBaseURL + envCfg.ValidateRefreshToken)

	requestBody := map[string]string{
		"user_id":       userId,
		"refresh_token": refreshToken,
	}

	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("failed to validate refresh token: %s", string(body))
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", nil
	}

	return tokenResponse.AccessToken, nil
}