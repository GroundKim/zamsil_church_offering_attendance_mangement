package controllers

import (
	"fmt"
	"strconv"
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
		models.GetAttendanceDiariesByYear(&attendanceDiaries, parsedYear)

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
		models.GetAttendanceDiariesByYear(&attendanceDiaries, parsedYear)

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

func GetAttendanceInfoByYear(c *gin.Context) {
	type attendanceInfo struct {
		Class    models.Class     `json:"class"`
		Teachers []models.Teacher `json:"teachers"`
		Students []models.Student `json:"students"`
	}

	type attendanceInfoByDate struct {
		Date            time.Time        `json:"date"`
		AttendanceInfos []attendanceInfo `json:"attendanceInfo"`
	}

	// get year from query
	year := ""
	if len(c.Query("year")) != 0 {
		year = c.Query("year")
	} else {
		c.JSON(400, gin.H{
			"err": "year from url paramemter is required",
		})
		return
	}

	parsedYear, _ := strconv.Atoi(year)
	// get class with year
	var classes []models.Class
	models.GetClassesWithYear(&classes, parsedYear)

	// get teacher with year
	var teachers []models.Teacher
	models.GetTeachersByYear(&teachers, parsedYear)

	// get attendance diaries with year
	var attendanceDiaries []models.AttendanceDiary
	timeParsedYear, _ := time.Parse("2006", year)
	models.GetAttendanceDiariesByYear(&attendanceDiaries, timeParsedYear)

	var attendanceInfoByDates []attendanceInfoByDate

	// fill attendance diaries
	for _, diary := range attendanceDiaries {
		isDateExisted := false
		for i, info := range attendanceInfoByDates {
			if info.Date == diary.AttendedAt {
				isDateExisted = true

				// check if the class is existed
				// if there is existed class
				isClassExisted := false
				for j, attendanceInfo := range info.AttendanceInfos {
					if attendanceInfo.Class.ID == diary.Student.ClassID {
						isClassExisted = true
						info.AttendanceInfos[j].Students = append(info.AttendanceInfos[j].Students, diary.Student)
					}
				}

				// no class existed, then push student with class
				if !isClassExisted {
					attendedStudent := diary.Student
					attendanceInfo := &attendanceInfo{Class: diary.Student.Class}
					attendanceInfo.Students = append(attendanceInfo.Students, attendedStudent)
					attendanceInfoByDates[i].AttendanceInfos = append(attendanceInfoByDates[i].AttendanceInfos, *attendanceInfo)
				}
			}
		}

		if !isDateExisted {
			// get teachers with class
			var infoTeachers []models.Teacher
			for _, teacher := range teachers {
				if teacher.ClassID == diary.Student.ClassID {
					infoTeachers = append(infoTeachers, teacher)
				}
			}

			// push teacher, class, and student
			attendanceInfo := attendanceInfo{Class: diary.Student.Class, Teachers: infoTeachers}
			attendanceInfo.Students = append(attendanceInfo.Students, diary.Student)
			attendanceInfoByDate := &attendanceInfoByDate{Date: diary.AttendedAt}
			attendanceInfoByDate.AttendanceInfos = append(attendanceInfoByDate.AttendanceInfos, attendanceInfo)
			attendanceInfoByDates = append(attendanceInfoByDates, *attendanceInfoByDate)
		}
	}

	// fill classes

	c.JSON(200, attendanceInfoByDates)
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

func GetAttendanceSummaryByYear(c *gin.Context) {
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

	parsedYear, _ := strconv.Atoi(year)
	timeParsedYear, _ := time.Parse("2006", year)
	var classes []models.Class
	models.GetClassesWithYear(&classes, parsedYear)

	var teachers []models.Teacher
	models.GetTeachersByYear(&teachers, parsedYear)

	var attendanceDiaries []models.AttendanceDiary
	models.GetAttendanceDiariesByYear(&attendanceDiaries, timeParsedYear)

	// json formatting
	type numberByDate struct {
		Date   string `json:"date"`
		Number int    `json:"number"`
	}

	type attendanceSummary struct {
		Class          models.Class     `json:"class"`
		ClassHeadcount int              `json:"classHeadcount"`
		Teachers       []models.Teacher `json:"teachers"`
		NumberByDates  []numberByDate   `json:"numberByDates"`
	}

	var attendanceSummaries []attendanceSummary

	// formatting by class
	for _, class := range classes {
		// get class headcount
		var student []models.Student
		models.GetStudentsByClass(&student, class)

		// initialize attendance summary with class and headcount
		attendanceSummary := &attendanceSummary{Class: class, ClassHeadcount: len(student)}

		// get teachers
		for _, teacher := range teachers {
			if class.ID == teacher.ClassID {
				attendanceSummary.Teachers = append(attendanceSummary.Teachers, teacher)
			}
		}

		// get number by attendance diaries with date
		for _, diary := range attendanceDiaries {
			if diary.Student.ClassID == class.ID {
				existed := false
				// check if existed attend at
				for i, numberByDate := range attendanceSummary.NumberByDates {
					if numberByDate.Date == diary.AttendedAt.Format("2006-01-02") {
						attendanceSummary.NumberByDates[i].Number++
						existed = true
						break
					}
				}

				if !existed {
					numberByDate := &numberByDate{Date: diary.AttendedAt.Format("2006-01-02"), Number: 1}
					attendanceSummary.NumberByDates = append(attendanceSummary.NumberByDates, *numberByDate)

				}
			}
		}

		attendanceSummaries = append(attendanceSummaries, *attendanceSummary)
	}

	c.JSON(200, attendanceSummaries)
}
