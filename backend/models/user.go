package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	ClientID string `gorm:"not null" json:"clientId"`
	Password string `gorm:"not null" json:"-"`
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
		HashPassword string
		ID           int
	}

	var result Result
	if err := DB.Table("user").Select("password", "id").Where("client_id = ?", user.ClientID).Scan(&result).Error; err != nil {
		return err
	}
	user.ID = result.ID

	if err := bcrypt.CompareHashAndPassword([]byte(result.HashPassword), []byte(user.Password)); err != nil {
		return err
	} else {
		return nil
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
