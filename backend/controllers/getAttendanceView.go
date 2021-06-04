package controllers

import (
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

type studentWithAttendance struct {
	Student models.Student       `json:"student"`
	Absence *models.AbsenceDiary `json:"absence"`
}

func GetAttendanceView(c *gin.Context) {
	date := c.Query("date")
	var studentsWithAttendance []studentWithAttendance
	if len(date) > 0 {
		parsedDate, _ := time.Parse("2006-01-02", date)

		var absenceDiaries []models.AbsenceDiary
		models.GetAbsenceDiariesByDate(&absenceDiaries, parsedDate)

		var attendanceDiaries []models.AttendanceDiary
		models.GetAttendanceDiariesByDate(&attendanceDiaries, parsedDate)

		var students []models.Student
		models.GetStudents(&students)

		// 1 means general absence
		var generalAbsence models.AbsenceDiary = models.AbsenceDiary{AbsenceTypeID: 1, AbsentAt: parsedDate}

		for _, student := range students {
			var studentWithAttendance studentWithAttendance = studentWithAttendance{Student: student, Absence: &generalAbsence}
			for _, absenceDiary := range absenceDiaries {
				// handling absent student
				if absenceDiary.StudentID == student.ID {
					studentWithAttendance.Absence = &absenceDiary
					continue
				}
			}

			for _, attendanceDiary := range attendanceDiaries {
				// handing absence student
				if attendanceDiary.StudentID == student.ID {
					studentWithAttendance.Absence = nil
				}
			}
			studentsWithAttendance = append(studentsWithAttendance, studentWithAttendance)
		}
	}

	c.JSON(200, studentsWithAttendance)
}

func GetAttendanceViewList(c *gin.Context) {
	// get list by Year
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
		models.GetAttendanceDiariesByDate(&attendanceDiaries, date)
		c.JSON(200, attendanceDiaries)
	}
}
