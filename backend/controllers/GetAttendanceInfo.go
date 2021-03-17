package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

type attendanceInfo struct {
	classID      int
	departmentID int
	teacherName  []string
	studentsName []string
}

// response with teacher, student , the number of classes
func GetAttendanceInfoForDepartment1(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	// get teacher info
	var teachers []models.Teacher
	models.GetTeacherByDeaprtment(&teachers, 1)

	// get student info
	var students []models.Student
	models.GetStudentByDepartment(&students, 1)

	var numberOfClass int64
	models.DB.Model(&models.Class{}).Where("department_id = 1").Count(&numberOfClass)

	var attendanceInfos []attendanceInfo
	var teachersName []string
	var studentsName []string

	// matching class, teacher, student
	for i := 0; i < int(numberOfClass); i++ {
		for j := 0; j < len(teachers); j++ {
			if teachers[j].ClassId == (i + 1) {
				teachersName = append(teachersName, teachers[j].Name)
				for k := 0; k < len(students); k++ {
					if students[k].ClassID == (i + 1) {
						studentsName = append(studentsName, students[k].Name)
					}
				}
			}
		}
		attendanceInfos = append(attendanceInfos, attendanceInfo{classID: (i + 1), departmentID: 1, teacherName: teachersName, studentsName: studentsName})
		teachersName = nil
		studentsName = nil
	}

	for _, v := range attendanceInfos {
		fmt.Println(v)
	}

}

func GetAttendanceInfoForDepartment2(c *gin.Context) {

	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	fmt.Println("1234")
	c.JSON(200, gin.H{
		"count":        15,
		"teacherName":  "선생님 예시",
		"studentNames": "김지상",
	})

}
