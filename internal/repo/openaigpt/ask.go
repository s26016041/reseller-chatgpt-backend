package openaigpt

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type AskInput struct {
	Model   string
	Message []openai.ChatCompletionMessage
	Tools   []openai.Tool
}

func (r *Repo) Ask(ctx context.Context, input *AskInput) (*openai.ChatCompletionMessage, error) {
	req := openai.ChatCompletionRequest{
		Model:    input.Model,
		Messages: input.Message,
		Tools:    input.Tools,
	}

	resp, err := r.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("CreateChatCompletion fail: %s", err.Error())
	}

	return &resp.Choices[0].Message, nil
}
