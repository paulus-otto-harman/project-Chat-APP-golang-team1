package handler

import (
	"go.uber.org/zap"
	"project/user-service/database"
	"project/user-service/service"
)

type Handler struct {
}

func NewHandler(service service.Service, logger *zap.Logger, rdb database.Cacher) *Handler {
	return &Handler{}
}
