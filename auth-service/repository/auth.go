package repository

import "gorm.io/gorm"

type Auth struct {
	Db *gorm.DB
}
