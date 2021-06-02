package controllers

import (
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetAttendanceView(c *gin.Context) {
	// get by Year
	if len(c.Query("year")) != 0 {
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
	// get by date
	if len(c.Query("date")) != 0 {
		date, _ := time.Parse("2006-01-02", c.Query("date"))

		var attendanceDiaries []models.AttendanceDiary
		models.GetAttendanceViewByDate(&attendanceDiaries, date)
		c.JSON(200, attendanceDiaries)
	}
}
