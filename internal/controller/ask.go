package controller

import (
	"fmt"
	"net/http"
	"reseller-chatgpt-backend/internal/env"
	"reseller-chatgpt-backend/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

type askInput struct{}

func (c *Controller) Ask(ctx *gin.Context) {
	input, err := newAskInput(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = c.Service.Ask(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"apikey": env.GetOpenAIAPIKey(),
	})
}

func newAskInput(ctx *gin.Context) (*askInput, error) {
	output := askInput{}

	err := utils.BindAll(ctx, &output)
	if err != nil {
		return nil, fmt.Errorf("BindAll fail: %s", err.Error())
	}

	return &output, nil
}
