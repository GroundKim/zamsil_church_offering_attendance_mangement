package models

import (
	"errors"
	"time"
	"zamsil_church_offering_attendance_mangement/config"

	"github.com/dgrijalva/jwt-go"
)

type AuthToken struct {
	ID    int
	Token string
}

func GenerateToken(conf *config.Config) (signedToken string, err error) {
	jwtClaim := &jwt.StandardClaims{
		ExpiresAt: time.Now().Local().AddDate(0, 3, 0).Unix(),
		Issuer:    "groundKim",
		IssuedAt:  time.Now().Local().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	if signedToken, err = token.SignedString([]byte(conf.AUTH.SECRETKEY)); err != nil {
		return
	}
	authToken := &AuthToken{Token: signedToken}
	// save the token in mariadb
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

	var authToken *AuthToken
	DB.Table("auth_token").Where("token = ?", signedToken).Find(&authToken)
	if authToken == nil {
		err = errors.New("token is not in DB")
		return
	}

	return
}

func (AuthToken) TableName() string {
	return "auth_token"
}
