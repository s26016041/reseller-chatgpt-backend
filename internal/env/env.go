package env

import (
	"os"
	"reseller-chatgpt-backend/internal/constant"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	if os.Getenv(constant.Localhost) == "true" {
		if err := godotenv.Load(); err != nil {
			panic("Error loading .env file")
		}
	}
}

func GetOpenAIAPIKey() string {
	return os.Getenv(constant.OpenAIAPIKey)
}

func GetCognitoLoginURL() string {
	return os.Getenv(constant.CognitoLoginURL)
}

func GetCognitoClientID() string {
	return os.Getenv(constant.CognitoClientID)
}

func GetSecretKey() string {
	return os.Getenv(constant.SecretKey)
}

func GetResellerURL() string {
	return os.Getenv(constant.ResellerURL)
}
