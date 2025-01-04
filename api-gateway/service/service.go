package service

import (
	"api_gateway/config"

	"go.uber.org/zap"
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
