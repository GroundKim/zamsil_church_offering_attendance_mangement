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
	// first validate with DB
	//make auth null
	AuthToken = &authToken{}
	DB.Preload("User").Where("token = ?", signedToken).Find(&AuthToken)
	if AuthToken.Token == "" {
		err = errors.New("token is not in DB")
		return
	}

	// validate with SECRET KEY
	claims = &jwt.StandardClaims{}
	if _, err = jwt.ParseWithClaims(
		signedToken,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(conf.AUTH.SECRETKEY), nil
		},
	); err != nil {
		err = errors.New("incorrect secret key")
		return
	}

	// validate with token expiration
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token is expired")
		return
	}

	return claims, err
}

func (authToken) TableName() string {
	return "auth_token"
}
