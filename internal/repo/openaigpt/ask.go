package openaigpt

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type AskInput struct {
	Model   string
	Message []openai.ChatCompletionMessage
}

func (r *Repo) Ask(ctx context.Context, input *AskInput) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:    input.Model,
		Messages: input.Message,
	}

	resp, err := r.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("CreateChatCompletion fail: %s", err.Error())
	}

	return resp.Choices[0].Message.Content, nil
}
