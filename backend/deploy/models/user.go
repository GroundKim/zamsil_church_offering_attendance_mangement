package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	ClientID string `json:"clientId"`
	Password string `json:"password"`
}

func (user *User) SaveUser() {
	// after hashing password, then create
	pwHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pwHash)
	DB.Create(&user)
}

func (user *User) ValidateUser() bool {
	var hashPassword string
	DB.Table("user").Select("password").Where("client_id = ?", user.ClientID).Scan(&hashPassword)

	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(user.Password)); err != nil {
		return false
	} else {
		return true
	}
}

func (User) TableName() string {
	return "user"
}
