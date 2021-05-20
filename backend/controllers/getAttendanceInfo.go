package controllers

import (
	"strconv"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

type attendanceInfo struct {
	DepartmentID      int                `json:"departmentId"`
	ClassName         string             `json:"className"`
	TeachersName      []string           `json:"teacherName"`
	StudentsIdAndName []studentIDAndName `json:"studentsIdandName"`
}

type studentIDAndName struct {
	StudentID   int    `json:"studentId"`
	StudentName string `json:"studentName"`
}

// response with teacher, student , the number of classes
func GetAttendanceInfoByDepartment(c *gin.Context) {

	departmentID, _ := strconv.Atoi(c.Query("department_id"))

	var classes []models.Class
	var teachers []models.Teacher
	var students []models.Student

	var attendanceInfos []attendanceInfo

	models.GetClassesByDepartment(&classes, departmentID)
	models.GetTeachersByDepartment(&teachers, departmentID)
	models.GetStudentsByDepartment(&students, departmentID)

	for _, class := range classes {
		attendance := &attendanceInfo{DepartmentID: departmentID, ClassName: class.Name}
		for _, teacher := range teachers {
			if class.ID == teacher.ClassID {
				attendance.TeachersName = append(attendance.TeachersName, teacher.Name)
			}
		}

		for _, student := range students {
			if class.ID == student.ClassID {
				idAndName := &studentIDAndName{student.ID, student.Name}
				attendance.StudentsIdAndName = append(attendance.StudentsIdAndName, *idAndName)
			}
		}

		attendanceInfos = append(attendanceInfos, *attendance)
	}

	c.JSON(200, attendanceInfos)

}
