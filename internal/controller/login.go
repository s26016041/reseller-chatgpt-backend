package controller

import (
	"fmt"
	"net/http"
	"reseller-chatgpt-backend/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

type loginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Controller) Login(ctx *gin.Context) {
	input, err := newLoginInput(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	output, err := c.Service.Login(ctx, input.Username, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": output})
}

func newLoginInput(ctx *gin.Context) (*loginInput, error) {
	output := loginInput{}

	err := utils.BindAll(ctx, &output)
	if err != nil {
		return nil, fmt.Errorf("BindAll fail: %s", err.Error())
	}

	return &output, nil
}
