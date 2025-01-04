package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"project/auth-service/model"
)

type OtpRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewOtpRepository(db *gorm.DB, log *zap.Logger) *OtpRepository {
	return &OtpRepository{db, log}
}

func (repo *OtpRepository) Create(otp *model.Otp) error {
	return repo.db.Create(&otp).Error
}

func (repo *OtpRepository) Get(criteria model.Otp) (model.Otp, error) {
	var otp model.Otp
	err := repo.db.Where(criteria).Preload("User").First(&otp).Error
	return otp, err
}
