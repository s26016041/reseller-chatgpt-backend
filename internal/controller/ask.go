package controller

import (
	"fmt"
	"net/http"
	"reseller-chatgpt-backend/internal/env"
	"reseller-chatgpt-backend/internal/pkg/utils"
	"reseller-chatgpt-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type askInput struct {
	Authorization string
	AskMessage    []askMessage `json:"askMessage"`
}

type askMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (c *Controller) Ask(ctx *gin.Context) {
	input, err := newAskInput(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	getOutput, err := c.Service.Ask(ctx, buildAskMessage(input.AskMessage))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"getOutput": getOutput,
	})
}

func newAskInput(ctx *gin.Context) (*askInput, error) {
	output := askInput{}

	err := utils.BindAll(ctx, &output)
	if err != nil {
		return nil, fmt.Errorf("BindAll fail: %s", err.Error())
	}

	output.Authorization = env.GetAuthorization(ctx)

	return &output, nil
}

func buildAskMessage(input []askMessage) *service.AskInput {
	output := service.AskInput{}

	for _, msg := range input {
		output.AskMessage = append(output.AskMessage, service.AskMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	return &output
}
