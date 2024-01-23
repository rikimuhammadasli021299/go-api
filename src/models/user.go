package models

import (
	"go-api/src/config"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email 	 string
	Password string
}

func CreateUser(newUser *User) *gorm.DB  {
	return config.DB.Create(&newUser)
}

func FindEmail(input *User) []User{
	// using raw query
	items := []User{}
	config.DB.Raw("SELECT * FROM users WHERE email = ?", input.Email).Scan(&items)
	return items

}
