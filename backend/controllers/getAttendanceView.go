package controllers

import (
	"fmt"
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
	year := c.Query("year")
	if len(date) > 0 {
		var studentsWithAttendance []studentWithAttendance
		parsedDate, _ := time.Parse("2006-01-02", date)

		var absenceDiaries []models.AbsenceDiary
		models.GetAbsenceDiariesByDate(&absenceDiaries, parsedDate)

		var attendanceDiaries []models.AttendanceDiary
		models.GetAttendanceDiariesByDate(&attendanceDiaries, parsedDate)

		var students []models.Student
		models.GetStudents(&students)

		// 1 means general absence
		var generalAbsence models.AbsenceDiary = models.AbsenceDiary{AbsenceType: models.AbsenceType{ID: 1, Name: "일반결석"}, AbsentAt: parsedDate}
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
		c.JSON(200, studentsWithAttendance)
		return
	}

	if len(year) > 0 {
		//get count by month and department
		var attendanceDiaries []models.AttendanceDiary
		parsedYear, _ := time.Parse("2006", year)
		models.GetAttendanceViewByYear(&attendanceDiaries, parsedYear)

		c.JSON(200, attendanceDiaries)
		return
	}

	c.JSON(400, gin.H{
		"err": "either year or date is required",
	})
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

func GetAttendanceNumberByMonth(c *gin.Context) {
	type attendanceNumberByMonth struct {
		Month                 int `json:"month"`
		NumberOfDepartmentOne int `json:"numberOfDepartmentOne"`
		NumberOfDepartmentTwo int `json:"numberOfDepartmentTwo"`
	}

	year := ""
	// get year from query
	if len(c.Query("year")) != 0 {
		year = c.Query("year")
	} else {
		c.JSON(400, gin.H{
			"err": "year from url paramemter is required",
		})
		return
	}

	//get count by month and department
	var attendanceNumberByMonths []attendanceNumberByMonth
	for i := 1; i < 13; i++ {
		var attendanceDiaries []models.AttendanceDiary
		date, _ := time.Parse("2006-01", fmt.Sprintf("%s-%02d", year, i))
		models.GetAttendanceDiariesByMonth(&attendanceDiaries, date)

		numberOfDepartmentOne := 0
		numberOfDepartmentTwo := 0

		for _, diary := range attendanceDiaries {
			departmentID := diary.Student.Class.DepartmentID
			if departmentID == 1 {
				numberOfDepartmentOne++
			} else if departmentID == 2 {
				numberOfDepartmentTwo++
			}
		}
		attendanceNumberByMonth := &attendanceNumberByMonth{Month: i, NumberOfDepartmentOne: numberOfDepartmentOne, NumberOfDepartmentTwo: numberOfDepartmentTwo}
		attendanceNumberByMonths = append(attendanceNumberByMonths, *attendanceNumberByMonth)
	}
	c.JSON(200, attendanceNumberByMonths)
}
