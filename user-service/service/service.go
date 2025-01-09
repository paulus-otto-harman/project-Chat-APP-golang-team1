package service

import (
	"project/user-service/repository"

	"go.uber.org/zap"
)

type Service struct {
	User *UserService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		User: NewUserService(repo, log),
	}
}
