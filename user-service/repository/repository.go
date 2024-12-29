package repository

import (
	"project/user-service/config"
	"project/user-service/database"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{}
}
