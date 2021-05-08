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
		var attendanceDiaries []models.AttendanceDiary
		models.GetAttendanceViewByYear(&attendanceDiaries, parsedYear)

		var attendedAts []time.Time
		for i := 0; i < len(attendanceDiaries); i++ {
			isFound := false
			for j := 0; j < len(attendedAts); j++ {
				if attendanceDiaries[i].AttendedAt.Format("2006-01-02") == attendedAts[j].Format("2006-01-02") {
					isFound = true
				}
			}

			if !isFound {
				attendedAts = append(attendedAts, attendanceDiaries[i].AttendedAt)
			}

		}

		payload := map[string][]time.Time{"attendedAts": attendedAts}
		c.JSON(200, payload)
	}
}
