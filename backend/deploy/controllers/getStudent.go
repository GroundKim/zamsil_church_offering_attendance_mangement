package controllers

import (
	"net/http"
	"strconv"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	if len(c.Query("department_id")) != 0 {
		departmentID, _ := strconv.Atoi(c.Query("department_id"))
		models.GetStudentsByDepartment(&students, departmentID)
		c.JSON(http.StatusOK, students)
		return
	}

	var studentsWithDepartment []models.StudentsWithDepartment
	models.GetStudentsWithDepartment(&studentsWithDepartment)
	c.JSON(http.StatusOK, studentsWithDepartment)
}
