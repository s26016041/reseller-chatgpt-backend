package service

import (
	"reseller-chatgpt-backend/internal/repo/openaigpt"
	"reseller-chatgpt-backend/internal/repo/resellerapi"
)

type Service struct {
	openaiGTPRepo   *openaigpt.Repo
	resellerAPIRepo *resellerapi.Repo
}

func NewService() *Service {
	return &Service{
		openaiGTPRepo:   openaigpt.NewRepo(),
		resellerAPIRepo: resellerapi.NewRepo(),
	}
}
