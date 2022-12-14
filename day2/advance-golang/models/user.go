package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
