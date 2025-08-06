package service

import (
	"context"
	"encoding/json"
	"fmt"
	"reseller-chatgpt-backend/internal/constant"
	"reseller-chatgpt-backend/internal/repo/openaigpt"

	"github.com/sashabaranov/go-openai"
)

type AskInput struct {
	Username   string
	Password   string
	AskMessage []AskMessage
}

type AskMessage struct {
	Role    string
	Content string
}

func (s *Service) Ask(ctx context.Context, input *AskInput) (string, error) {

	askInput, err := getAskInput(input)
	if err != nil {
		return "", fmt.Errorf("getAskInput fail: %s", err.Error())
	}

	output, err := s.openaiGTPRepo.Ask(ctx, askInput)
	if err != nil {
		return "", fmt.Errorf("Ask fail: %s", err.Error())
	}

	if len(output.ToolCalls) > 0 {
		err = s.funcCall(askInput, output, input.Username, input.Password)
		if err != nil {
			return "", fmt.Errorf("funcCall fail: %s", err.Error())
		}

		output, err = s.openaiGTPRepo.Ask(ctx, askInput)
		if err != nil {
			return "", fmt.Errorf("Ask fail: %s", err.Error())
		}
	}

	return output.Content, nil
}

// openai.ChatMessageRoleUser      // "user"
// openai.ChatMessageRoleSystem    // "system"
// openai.ChatMessageRoleAssistant // "assistant"
// openai.ChatMessageRoleTool      // "tool"

const systemContent = `你是一位專業客服人員，負責協助回答顧客對產品、訂單、付款等相關問題。
請使用親切、清楚、具體的語氣回應客戶問題。
當客戶的問題需要查詢資料時，請使用工具 API 來查詢。
如果無法解答，也要禮貌告知並建議客戶聯絡人工客服。請用客人所使用語言回答。`

func getAskInput(input *AskInput) (*openaigpt.AskInput, error) {
	tools, err := getTool()
	if err != nil {
		return nil, fmt.Errorf("getFunctionCall fail: %s", err.Error())
	}

	output := openaigpt.AskInput{
		Model: openai.GPT4Dot1,
		Message: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemContent,
			},
		},
		Tools: tools,
	}

	for _, msg := range input.AskMessage {
		output.Message = append(output.Message, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	return &output, nil
}

func getTool() ([]openai.Tool, error) {
	functionLicenseInventory, err := getFunctionLicenseInventory()
	if err != nil {
		return nil, fmt.Errorf("getFunctionLicenseInventory fail: %s", err.Error())
	}

	output := []openai.Tool{
		{
			Type:     openai.ToolTypeFunction,
			Function: functionLicenseInventory,
		},
	}

	return output, nil
}

const getLicenseInventoryDescription = "取得目前可用的 license 庫存數量"

func getFunctionLicenseInventory() (*openai.FunctionDefinition, error) {
	output := openai.FunctionDefinition{
		Name:        constant.FuncLicenseInventory,
		Description: getLicenseInventoryDescription,
		Parameters: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	}
	return &output, nil
}

func (s *Service) funcCall(askInput *openaigpt.AskInput, message *openai.ChatCompletionMessage, username string, password string) error {
	askInput.Message = append(askInput.Message, openai.ChatCompletionMessage{
		Role:       message.Role,
		ToolCalls:  message.ToolCalls,
		ToolCallID: message.ToolCallID,
		Content:    message.Content,
	})

	for _, tool := range message.ToolCalls {
		json, err := s.funcCallSwitch(tool.Function.Name, username, password)
		if err != nil {
			return fmt.Errorf("funcCallSwitch fail: %s", err.Error())
		}

		askInput.Message = append(askInput.Message, openai.ChatCompletionMessage{
			Role:       openai.ChatMessageRoleTool,
			ToolCallID: tool.ID,
			Content:    json + "remaining 代表剩餘天數，請將天數轉換為「X 年 Y 個月 Z 天」的格式，其中：  - 30 天視為 1 個月  - 360 天視為 1 年",
		})
	}

	return nil
}

func (s *Service) funcCallSwitch(funcName string, username string, password string) (string, error) {
	switch funcName {
	case constant.FuncLicenseInventory:
		return s.getlicensesInventory(username, password)
	default:
		return "", fmt.Errorf("unknown function call: %s", funcName)
	}
}

func (s *Service) getlicensesInventory(username string, password string) (string, error) {
	authorization, err := s.resellerAPIRepo.Login(username, password)
	if err != nil {
		return "", fmt.Errorf("Login fail: %s", err.Error())
	}

	info, err := s.resellerAPIRepo.Info(authorization)
	if err != nil {
		return "", fmt.Errorf("info fail: %s", err.Error())
	}

	inventory, err := s.resellerAPIRepo.LicensesInventory(authorization, info.ID)
	if err != nil {
		return "", fmt.Errorf("info fail: %s", err.Error())
	}

	jsonBytes, err := json.Marshal(inventory)
	if err != nil {
		return "", fmt.Errorf("json.Marshal fail: %s", err.Error())
	}

	return string(jsonBytes), nil
}
