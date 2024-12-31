package model

import "gorm.io/gorm"

type User struct {
	Username string
	Password string
	gorm.Model
}
