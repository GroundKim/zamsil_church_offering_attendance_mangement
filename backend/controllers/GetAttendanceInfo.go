package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

type attendanceInfo struct {
	ClassID      int      `json:"classId"`
	DepartmentID int      `json:"departmentId"`
	TeacherName  []string `json:"teacherName"`
	StudentsName []string `json:"studentName"`
}

// response with teacher, student , the number of classes
func GetAttendanceInfoByDepartment(department int) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

		// get teacher info
		var teachers []models.Teacher
		models.GetTeacherByDeaprtment(&teachers, 1)

		// get student info
		var students []models.Student
		models.GetStudentByDepartment(&students, 1)

		var numberOfClass int64
		models.DB.Model(&models.Class{}).Where(fmt.Sprintf("department_id = %d", department)).Count(&numberOfClass)

		var attendanceInfos []attendanceInfo
		var teachersName []string
		var studentsName []string

		// matching class, teacher, student
		for classId := 1; classId <= int(numberOfClass); classId++ {
			for j := 0; j < len(teachers); j++ {
				if teachers[j].ClassID == classId {
					teachersName = append(teachersName, teachers[j].Name)
				}
			}

			for j := 0; j < len(students); j++ {
				if students[j].ClassID == classId {
					studentsName = append(studentsName, students[j].Name)
				}
			}

			attendanceInfos = append(attendanceInfos, attendanceInfo{ClassID: classId, DepartmentID: 1, TeacherName: teachersName, StudentsName: studentsName})
			teachersName = nil
			studentsName = nil

		}
		c.JSON(200, attendanceInfos)
	}
	return gin.HandlerFunc(fn)
}
