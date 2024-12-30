package repository

import (
	"go.uber.org/zap"
	"project/api-gateway/config"
	"project/api-gateway/database"
)

type Repository struct {
}

func NewRepository(cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{}
}
