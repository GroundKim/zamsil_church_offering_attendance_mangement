package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

// save attendance information into database
func PostAttendanceInfo(c *gin.Context) {
	var attendanceDiaries []models.AttendanceDiary
	if err := c.Bind(&attendanceDiaries); err != nil {
		fmt.Println("Error in json bind: post attendance diary ", err)
	}

	for _, attendanceDiary := range attendanceDiaries {
		attendanceDiary.SaveAttendanceDiary()
	}
}
