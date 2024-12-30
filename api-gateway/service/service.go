package service

import (
	"project/api-gateway/config"
	"project/api-gateway/repository"

	"go.uber.org/zap"
)

type Service struct {
}

func NewService(repo repository.Repository, appConfig config.Config, log *zap.Logger) Service {
	return Service{}
}
