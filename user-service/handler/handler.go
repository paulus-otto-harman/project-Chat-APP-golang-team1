package handler

import (
	"user_service/database"
	"user_service/service"

	"go.uber.org/zap"
)

type Handler struct {
}

func NewHandler(service service.Service, logger *zap.Logger, rdb database.Cacher) *Handler {
	return &Handler{}
}
