package service

import (
	"context"
	"fmt"
	"reseller-chatgpt-backend/internal/repo/openaigpt"

	"github.com/sashabaranov/go-openai"
)

type AskInput struct {
	AskMessage []AskMessage
}

type AskMessage struct {
	Role    string
	Content string
}

func (s *Service) Ask(ctx context.Context, input *AskInput) (string, error) {
	output, err := s.openaiGTPRepo.Ask(ctx, getAskInput(input))
	if err != nil {
		return "", fmt.Errorf("Ask fail: %s", err.Error())
	}

	return output, nil
}

// openai.ChatMessageRoleUser      // "user"
// openai.ChatMessageRoleSystem    // "system"
// openai.ChatMessageRoleAssistant // "assistant"
// openai.ChatMessageRoleTool      // "tool"

const systemContent = `你是一位專業客服人員，負責協助回答顧客對產品、訂單、付款等相關問題。
請使用親切、清楚、具體的語氣回應客戶問題。
當客戶的問題需要查詢資料時，請使用工具 API 來查詢。
如果無法解答，也要禮貌告知並建議客戶聯絡人工客服。請用客人所使用語言回答。`

func getAskInput(input *AskInput) *openaigpt.AskInput {
	output := openaigpt.AskInput{
		Model: openai.GPT4Dot1,
		Message: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemContent,
			},
		},
	}

	for _, msg := range input.AskMessage {
		output.Message = append(output.Message, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	return &output
}
