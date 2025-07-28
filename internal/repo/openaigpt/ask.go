package openaigpt

import (
	"context"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
)

type AskInput struct {
	Model   string
	Message []openai.ChatCompletionMessage
}

func (r *Repo) Ask(ctx context.Context, input *AskInput) error {
	req := openai.ChatCompletionRequest{
		Model:    input.Model,
		Messages: input.Message,
	}

	resp, err := r.Client.CreateChatCompletion(ctx, req)
	if err != nil {
		log.Fatalf("OpenAI API fail: %s", err.Error())
	}

	fmt.Println("ChatGPT 回覆：")
	fmt.Println(resp.Choices[0].Message.Content)

	return nil
}
