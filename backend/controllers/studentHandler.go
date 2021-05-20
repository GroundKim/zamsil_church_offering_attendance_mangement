package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func PostStudent(c *gin.Context) {

	var NewStudents []models.Student
	if err := c.BindJSON(&NewStudents); err != nil {
		fmt.Println("Error in json bind: post specific offering diary", err)
	}
	models.SaveStudents(NewStudents)

}
