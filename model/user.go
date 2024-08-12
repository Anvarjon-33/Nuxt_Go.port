package model

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type AuthUser struct {
	//gorm.Model
	ID       uuid.UUID `gorm:"id" json:"id" form:"id"`
	Name     string    `gorm:"username" json:"username" form:"username"`
	Email    string    `gorm:"email;unique;size:100" json:"email" form:"email"`
	Password string    `gorm:"password" json:"password" form:"password"`
	Token    string    `gorm:"access_token" json:"auth.token" form:"password"`
}

type Auth struct {
}

var DB, err = gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})

func init() {
	err := DB.AutoMigrate(&AuthUser{})
	if err != nil {
		panic(err)
	}
}
