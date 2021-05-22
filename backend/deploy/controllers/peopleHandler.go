package controllers

import (
	"fmt"
	"net/http"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func SaveStudents(c *gin.Context) {
	var NewStudents []models.Student
	var err error
	if err = c.BindJSON(&NewStudents); err != nil {
		fmt.Println("Error in json bind: post specific offering diary", err)
	}
	if err = models.SaveStudents(NewStudents); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error in SaveStudents()",
		})
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "save completely",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
}

func PatchStudent(c *gin.Context) {
	var student models.Student
	var err error
	if err = c.BindJSON(&student); err != nil {
		fmt.Println("Error in json bind: PatchStudent", err)
	}

}

func SaveTeachers(c *gin.Context) {
	var NewTeacher []models.Teacher
	var err error
	if err = c.BindJSON(&NewTeacher); err != nil {
		fmt.Println("Error in json bind: post specific offering diary", err)
	}
	if err = models.SaveTeachers(NewTeacher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error in SaveTeacher()",
		})
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "save completely",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
}
