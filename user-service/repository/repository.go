package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	User RepositoryUser
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		User: NewRepositoryUser(db),
	}
}
