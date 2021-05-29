package controllers

import (
	"fmt"
	"net/http"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

type classInfo struct {
	Class    models.Class
	Students []models.Student // `json:"name"`
	Teachers []models.Teacher
}

func GetMembers(c *gin.Context) {
	var classInfos []classInfo

	var classes []models.Class
	models.GetClasses(&classes)

	for _, class := range classes {
		departmentName := class.Department.DepartmentName

		var students []models.Student
		models.GetStudentByClassNameAndDepartment(&students, departmentName, class.Name)

		var teachers []models.Teacher
		models.GetTeacherByClassNameAndDepartment(&teachers, departmentName, class.Name)

		var classInfo = classInfo{class, students, teachers}
		classInfos = append(classInfos, classInfo)
	}

	c.JSON(http.StatusOK, classInfos)
}

func SaveStudents(c *gin.Context) {
	var NewStudents []models.Student
	var err error
	if err = c.BindJSON(&NewStudents); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{
			"error": "Error in JSON Binding",
		})
		fmt.Println(err)
		return
	}
	if err = models.SaveStudents(&NewStudents); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error ins saving students",
		})
		return
	}

	if err == nil {
		c.JSON(http.StatusOK, NewStudents)
		return
	}
}

func DeleteStudents(c *gin.Context) {
	var students []models.Student
	var err error
	if err = c.BindJSON(&students); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{
			"error": "Error in JSON Binding",
		})
		return
	}

	if err = models.DeleteStudents(&students); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in delete students",
		})
		return
	}

	if err == nil {
		c.JSON(http.StatusOK, students)
		return
	}
}

func SaveTeachers(c *gin.Context) {
	var NewTeacher []models.Teacher
	var err error
	if err = c.BindJSON(&NewTeacher); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{
			"error": "Error in JSON Binding",
		})
		return
	}
	if err = models.SaveTeachers(NewTeacher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error in SaveTeacher()",
		})
		return
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "save completely",
		})
	}
}

func PutStudent(c *gin.Context) {
	var student models.Student
	var err error
	if err = c.BindJSON(&student); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{
			"error": "Error in JSON Binding",
		})
		fmt.Println(err)
		return
	}

	if err = models.PutStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error in SaveTeacher()",
		})
		return
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "put completely",
		})
	}

}
