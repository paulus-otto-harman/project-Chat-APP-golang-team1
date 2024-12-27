package handler

import (
	"go.uber.org/zap"
	"project/auth-service/database"
	"project/auth-service/service"
)

type Handler struct {
}

func NewHandler(service service.Service, logger *zap.Logger, rdb database.Cacher) *Handler {
	return &Handler{}
}
