package service

import (
	"go.uber.org/zap"
	"project/api-gateway/config"
)

type Service struct {
	Auth  AuthService
	Email EmailService
}

func NewService(appConfig config.Config, log *zap.Logger) Service {
	return Service{
		Auth:  NewAuthService(appConfig.MicroserviceConfig.Auth, log),
		Email: NewEmailService(appConfig.Email, log),
	}
}
