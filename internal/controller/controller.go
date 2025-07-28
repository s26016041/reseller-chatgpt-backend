package controller

import "reseller-chatgpt-backend/internal/service"

func NewController() *Controller {
	return &Controller{
		Service: service.NewService(),
	}
}

type Controller struct {
	Service *service.Service
}
