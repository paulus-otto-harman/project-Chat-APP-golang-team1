package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"project/auth-service/model"
)

type AuthRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewAuthRepository(db *gorm.DB, log *zap.Logger) *AuthRepository {
	return &AuthRepository{db: db, log: log}
}

func (repo *AuthRepository) Create(user *model.User) error {
	return repo.db.Create(user).Error
}
