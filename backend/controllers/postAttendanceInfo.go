package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

// save attendance information into database
func PostAttendanceInfo(c *gin.Context) {

	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	var attendanceDiarys []models.AttendanceDiary
	if err := c.Bind(&attendanceDiarys); err != nil {
		fmt.Println("Error in json bind ", err)
	}

	for _, attendanceDiary := range attendanceDiarys {
		attendanceDiary.SaveAttendanceDiary()
	}
}
