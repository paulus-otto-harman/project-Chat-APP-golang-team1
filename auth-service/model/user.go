package model

import "gorm.io/gorm"

type User struct {
	Email string
	gorm.Model
}
