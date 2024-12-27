package repository

import (
	"project/auth-service/config"
	"project/auth-service/database"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{}
}
