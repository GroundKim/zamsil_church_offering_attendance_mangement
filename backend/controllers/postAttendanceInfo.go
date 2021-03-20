package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

// save attendance information into database
func PostAttendanceInfo() gin.HandlerFunc {
	fn := func(c *gin.Context) {

		var attendanceDiarys []models.AttendanceDiary
		if err := c.Bind(&attendanceDiarys); err != nil {
			fmt.Println("Error in json bind")
		}

		for _, attendanceDiary := range attendanceDiarys {
			attendanceDiary.SaveAttendanceDiary()
		}
	}
	return gin.HandlerFunc(fn)
}
