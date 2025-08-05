package openaigpt

import (
	"reseller-chatgpt-backend/internal/env"

	"github.com/sashabaranov/go-openai"
)

type Repo struct {
	Client *openai.Client
}

func NewRepo() *Repo {
	return &Repo{
		Client: openai.NewClient(env.GetOpenAIAPIKey()),
	}
}
