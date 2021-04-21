package controllers

import (
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetAttendanceView(c *gin.Context) {
	if len(c.Query("year")) != 0 {

		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

		year := c.Query("year")
		parsedYear, _ := time.Parse("2006", year)
		var attendanceDiarys []models.AttendanceDiary
		var attendedAts []time.Time
		models.GetAttendanceViewByYear(&attendanceDiarys, parsedYear)

		for i := 0; i < len(attendedAts); i++ {
			isFound := false
			for j := i; j < len(attendedAts)-i; j++ {

				if attendanceDiarys[i].AttendedAt.Format("2006-01-02") == attendanceDiarys[j].AttendedAt.Format("2006-01-02") {
					isFound = true
				}
			}

			if !isFound {
				attendedAts = append(attendedAts, attendanceDiarys[i].AttendedAt)
			}

		}
		payload := map[string][]time.Time{"attendedAts": attendedAts}
		c.JSON(200, payload)
	}
}
