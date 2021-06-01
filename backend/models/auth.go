package models

import (
	"errors"
	"time"
	"zamsil_church_offering_attendance_mangement/config"

	"github.com/dgrijalva/jwt-go"
)

type authToken struct {
	ID     int
	UserID int
	Token  string

	User User `gorm:"references:ID"`
}

var AuthToken *authToken

func GenerateToken(conf *config.Config, user User) (signedToken string, err error) {
	jwtClaim := &jwt.StandardClaims{
		ExpiresAt: time.Now().Local().AddDate(0, 3, 0).Unix(),
		Issuer:    "groundKim",
		IssuedAt:  time.Now().Local().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	if signedToken, err = token.SignedString([]byte(conf.AUTH.SECRETKEY)); err != nil {
		return
	}
	authToken := &authToken{UserID: user.ID, Token: signedToken}

	// save the token with UserID in db
	DB.Create(&authToken)
	return
}

func ValidateToken(signedToken string, conf *config.Config) (claims *jwt.StandardClaims, err error) {
	claims = &jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(
		signedToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.AUTH.SECRETKEY), nil
		},
	)

	if err != nil {
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return
	}

	// it reqires two select queries, which doesn't look so good.
	DB.Preload("User").Where("token = ?", signedToken).Find(&AuthToken)
	// Tried to join raw query in order to query once, but it doesn't save well.
	// DB.Raw("SELECT auth_token.* FROM auth_token INNER JOIN user on auth_token.user_id = user.id WHERE token = ?", signedToken).Find(&authToken)
	if AuthToken == nil {
		err = errors.New("token is not in DB")
		return
	}

	return
}

func (authToken) TableName() string {
	return "auth_token"
}
