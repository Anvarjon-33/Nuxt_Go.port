package model

import (
	"Anvarjon-33/Nuxt_Go/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthUser struct {
	gorm.Model
	ID       uuid.UUID `gorm:"id" json:"id" form:"id"`
	Name     string    `gorm:"username" json:"username" form:"username"`
	Email    string    `gorm:"email;unique;size:100" json:"email" form:"email"`
	Password string    `gorm:"password" json:"password" form:"password"`
	Token    string    `gorm:"access_token" json:"auth.token" form:"password"`
}

type Auth struct{}

func init() {
	if db.DB != nil {
		err := db.DB.AutoMigrate(&AuthUser{})
		if err != nil {
		}
	}
}
