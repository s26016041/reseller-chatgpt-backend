package service

import "reseller-chatgpt-backend/internal/repo/openaigpt"

type Service struct {
	openaiGTPRepo *openaigpt.Repo
}

func NewService() *Service {
	return &Service{
		openaiGTPRepo: openaigpt.NewOpenAIGPT(),
	}
}
