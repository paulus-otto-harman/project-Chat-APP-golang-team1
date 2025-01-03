package handler

import (
	"project/chat-service/service"

	"go.uber.org/zap"
)

type Handler struct {
	ChatHandler ChatHandler
}

func NewHandler(service service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		ChatHandler: *NewChatHandler(service, logger),
	}
}
