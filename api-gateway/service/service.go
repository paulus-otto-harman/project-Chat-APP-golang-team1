package service

import (
	"go.uber.org/zap"
	"project/api-gateway/config"
)

type Service struct {
	Auth AuthService
}

func NewService(appConfig config.Config, log *zap.Logger) Service {
	return Service{
		Auth: NewAuthService(log),
	}
}
