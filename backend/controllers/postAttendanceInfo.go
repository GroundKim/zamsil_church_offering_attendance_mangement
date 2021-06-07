package controllers

import (
	"fmt"
	"net/http"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

// save attendance information into database
func PostAttendanceDiary(c *gin.Context) {
	var attendanceDiaries []models.AttendanceDiary
	if err := c.Bind(&attendanceDiaries); err != nil {
		fmt.Println("Error in json bind: post attendance diary ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "fail in binding attendanceDiary",
		})
		return
	}

	for _, attendanceDiary := range attendanceDiaries {
		if err := attendanceDiary.SaveAttendanceDiary(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "fail in save attendance Diary",
			})
			return
		}
	}

	// log
	models.AuthToken.User.ActiveStamp("Save attendanceDiary", "")
}
