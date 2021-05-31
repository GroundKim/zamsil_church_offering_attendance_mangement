package middlewares

import (
	"fmt"
	"net/http"
	"zamsil_church_offering_attendance_mangement/config"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func Authorize(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken, err := c.Request.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusForbidden, "No Authorization token")
			c.Abort()
			return
		}
		token := clientToken.Value

		_, err = models.ValidateToken(token, conf)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "token Unauthorized")
			c.Abort()
			fmt.Println(err)
			return
		}

		// write down log
		user := models.GetUserByToken(token)
		action := c.Request.Method
		clientIP := c.ClientIP()
		URL := c.Request.URL.String()
		user.ActiveStamp(clientIP, action, URL)

		c.Next()
	}

}
