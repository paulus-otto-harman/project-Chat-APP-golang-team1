package service

import (
	"project/auth-service/config"
	"project/auth-service/repository"

	"go.uber.org/zap"
)

type Service struct {
}

func NewService(repo repository.Repository, appConfig config.Config, log *zap.Logger) Service {
	return Service{}
}
