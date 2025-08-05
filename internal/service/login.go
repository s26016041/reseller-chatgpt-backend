package service

import (
	"context"
	"fmt"
	"reseller-chatgpt-backend/internal/pkg/utils"
)

func (s *Service) Login(ctx context.Context, username string, password string) (string, error) {
	_, err := s.resellerAPIRepo.Login(username, password)
	if err != nil {
		return "", fmt.Errorf("Login fail: %s", err.Error())
	}

	token, err := utils.GenerateJWT(username, password)
	if err != nil {
		return "", fmt.Errorf("GenerateJWT fail: %s", err.Error())
	}

	return token, nil
}
