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

		c.Next()
	}

}
