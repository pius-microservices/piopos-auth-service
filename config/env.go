package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	BaseURL string

	UserServiceBaseURL string
	GetUserByEmail     string
	CreateRefreshToken string

	JwtSecret string
	Mode      string
}

var (
	envConfig *Config
	once      sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		}

		envConfig = &Config{
			Port:               os.Getenv("APP_PORT"),
			BaseURL:            os.Getenv("BASE_URL"),

			UserServiceBaseURL: os.Getenv("USER_SERVICE_BASE_URL"),
			GetUserByEmail:     os.Getenv("GET_USER_BY_EMAIL"),
			CreateRefreshToken: os.Getenv("CREATE_REFRESH_TOKEN"),
			
			JwtSecret:          os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
			Mode:               os.Getenv("MODE"),
		}
	})

	return envConfig
}
