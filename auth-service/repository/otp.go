package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OtpRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewOtpRepository(db *gorm.DB, log *zap.Logger) *OtpRepository {
	return &OtpRepository{db, log}
}

func (s *OtpRepository) Create() {}
