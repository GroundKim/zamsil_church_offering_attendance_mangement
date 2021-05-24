package controllers

import (
	"fmt"
	"net/http"
	"zamsil_church_offering_attendance_mangement/config"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func Login(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			fmt.Println("Error in login user binding:", err)
			return
		}

		if user.ValidateUser() {
			token, _ := models.GenerateToken(conf)
			domain := conf.COOKIE.DOMAIN
			c.SetCookie("auth_token", token, 60*60*24*31*3, "/", domain, conf.COOKIE.SECURE, conf.COOKIE.HTTPONLY)

			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})

		} else {
			c.JSON(http.StatusUnauthorized, "Enter correct ID with password")
			return
		}
	}
}
