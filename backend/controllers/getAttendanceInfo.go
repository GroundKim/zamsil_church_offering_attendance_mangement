package controllers

import (
	"strconv"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

type attendanceInfo struct {
	ClassName         int                `json:"classId"`
	DepartmentID      int                `json:"departmentId"`
	TeachersName      []string           `json:"teacherName"`
	StudentsIDandName []studentIDandName `json:"studentsIdandName"`
}

type studentIDandName struct {
	StudentID   int    `json:"studentId"`
	StudentName string `json:"studentName"`
}

// response with teacher, student , the number of classes
func GetAttendanceInfoByDepartment(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	departmentID, _ := strconv.Atoi(c.Query("department_id"))

	// get teacher info
	var teachers []models.Teacher
	models.GetTeacherByDeaprtment(&teachers, departmentID)

	// get student info
	var students []models.Student
	models.GetStudentByDepartment(&students, departmentID)

	// var numberOfClass int64
	// models.DB.Model(&models.Class{}).Where(fmt.Sprintf("department_id = %d", departmentID)).Count(&numberOfClass)

	var classes []models.Class
	models.GetClasses(&classes)

	var attendanceInfos []attendanceInfo
	var teachersName []string
	var studentsIDandName []studentIDandName

	// matching class, teacher, student
	for classId := 1; classId <= int(numberOfClass); classId++ {
		for j := 0; j < len(teachers); j++ {
			if teachers[j].ClassID == classId {
				teachersName = append(teachersName, teachers[j].Name)
			}
		}

		for j := 0; j < len(students); j++ {
			if students[j].ClassID == classId {
				studentsIDandName = append(studentsIDandName, studentIDandName{StudentID: students[j].ID, StudentName: students[j].Name})
			}
		}

		attendanceInfos = append(attendanceInfos, attendanceInfo{ClassID: classId, DepartmentID: 1, TeachersName: teachersName, StudentsIDandName: studentsIDandName})
		teachersName = nil
		studentsIDandName = nil

	}
	c.JSON(200, attendanceInfos)

}
