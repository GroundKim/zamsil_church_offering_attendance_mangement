package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	ClientID string `gorm:"not null" json:"clientId"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"userName"`
	Role     string `gorm:"not null" json:"userStatus"`
}

type Log struct {
	ID     int
	UserID int       `gorm:"not null"`
	IP     string    `gorm:"not null"`
	Date   time.Time `gorm:"not null"`
	Action string    `gorm:"not null"`
	Note   string    `gorm:"type:longtext; default: null"`

	User User `gorm:"references:ID"`
}

var ClientIP string

func (user *User) SaveUser() {
	// after hashing password, then create
	pwHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pwHash)
	DB.Create(&user)
}

func (user *User) ValidateUser() error {
	type Result struct {
		ID       int
		Password string
	}

	// get valid id and pw with client ID
	var result Result
	if err := DB.Table("user").Select("id", "password").Where("client_id = ?", user.ClientID).Scan(&result).Error; err != nil {
		return err
	}

	// validate password
	if result.ID != 0 {
		user.ID = result.ID

		if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
			return err
		} else {
			return nil
		}

	} else {
		err := errors.New("info from client is invalid")
		return err
	}

}

func (user *User) LoginStamp(IP string) {
	log := &Log{UserID: user.ID, IP: IP, Date: time.Now(), Action: "login"}
	DB.Create(&log)

}

func (user *User) ActiveStamp(active string, note string) {
	log := &Log{UserID: user.ID, IP: ClientIP, Date: time.Now(), Action: active, Note: note}
	DB.Create(&log)
}

func GetUserByToken(token string) User {
	var user User
	DB.Raw("SELECT user.* FROM user INNER JOIN auth_token ON user.id = auth_token.user_id WHERE auth_token.token = ?", token).Find(&user)
	return user
}

func (User) TableName() string {
	return "user"
}

func (Log) TableName() string {
	return "log"
}
