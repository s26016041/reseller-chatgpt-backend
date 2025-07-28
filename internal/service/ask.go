package service

import (
	"context"
	"fmt"
	"reseller-chatgpt-backend/internal/repo/openaigpt"

	"github.com/sashabaranov/go-openai"
)

func (s *Service) Ask(ctx context.Context) error {
	err := s.openaiGTPRepo.Ask(ctx, getAskInput())
	if err != nil {
		return fmt.Errorf("Ask fail: %s", err.Error())
	}

	return nil
}

func getAskInput() *openaigpt.AskInput {
	output := openaigpt.AskInput{
		Model: openai.GPT4Dot1,
		Message: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "用繁體中文幫我解釋什麼是 middleware",
			},
		},
	}

	return &output
}
