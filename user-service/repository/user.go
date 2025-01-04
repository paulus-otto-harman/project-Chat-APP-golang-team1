package repository

import (
	"errors"
	"log"
	"user_service/model"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	GetAllUsers(filter bool) ([]model.User, error)
	Insert(user *model.User) error
	UpdateProfile(user *model.User) error
}

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) RepositoryUser {
	return &repositoryUser{db}
}

func (r *repositoryUser) GetAllUsers(filter bool) ([]model.User, error) {
	var users []model.User
	if filter {
		err := r.db.Where("is_online=?", filter).Find(&users).Error
		if err != nil {
			log.Println(err)
			return []model.User{}, errors.New(" Internal Server Error")
		}
	} else {
		if err := r.db.Find(&users).Error; err != nil {
			log.Println(err)
			return []model.User{}, errors.New(" Internal Server Error")
		}
	}
	return users, nil
}
func (repo *repositoryUser) Insert(user *model.User) error {
	err := repo.db.Create(&user).Error
	if err != nil {
		log.Println(err)
		return errors.New(" Already Registered")
	}
	return nil
}
func (repo *repositoryUser) UpdateProfile(user *model.User) error {
	err := repo.db.Model(&model.User{}).
		Where("email ILIKE ?", user.Email).
		Updates(user).Error
	if err != nil {
		log.Println(err)
		return errors.New(" Bad Request")
	}
	return nil
}
