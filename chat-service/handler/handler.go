package handler

import (
	"go.uber.org/zap"
	"project/chat-service/database"
	"project/chat-service/service"
)

type Handler struct {
}

func NewHandler(service service.Service, logger *zap.Logger, rdb database.Cacher) *Handler {
	return &Handler{}
}
