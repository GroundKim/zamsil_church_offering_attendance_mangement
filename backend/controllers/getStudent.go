package controllers

import (
	"strconv"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	if len(c.Query("department_id")) != 0 {
		departmentID, _ := strconv.Atoi(c.Query("department_id"))
		models.GetStudentByDepartment(&students, departmentID)
		c.JSON(200, students)
		return
	}

	models.GetStudents(&students)
	c.JSON(200, students)
}
