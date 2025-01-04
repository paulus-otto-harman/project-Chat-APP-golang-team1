package service

import (
	"project/user-service/config"
	"project/user-service/repository"

	"go.uber.org/zap"
)

type Service struct {
	User *ServiceUser
}

func NewService(repo repository.Repository, appConfig config.Config, log *zap.Logger) Service {
	return Service{
		User: NewServiceUser(repo, log),
	}
}
