package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

type classInfo struct {
	Class    models.Class     `json:"class"`
	Students []models.Student `json:"students"`
	Teachers []models.Teacher `json:"teachers"`
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
			"error": "Error in saving students",
		})
		return
	}

	if err == nil {
		c.JSON(http.StatusOK, NewStudents)
		return
	}
}

func DeleteStudent(c *gin.Context) {
	var err error
	var student models.Student
	studentId := c.Param("id")
	student.ID, _ = strconv.Atoi(studentId)
	if err = models.DeleteStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in delete students",
		})
		return
	}

	if err == nil {
		c.JSON(http.StatusOK, student)
		return
	}
}

func GetTeachers(c *gin.Context) {
	var teachers []models.Teacher
	models.GetTeachers(&teachers)

	c.JSON(http.StatusOK, teachers)
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
	studentID := c.Param("id")
	var err error
	if err = c.BindJSON(&student); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{
			"error": "Error in JSON Binding",
		})
		fmt.Println(err)
		return
	}

	student.ID, _ = strconv.Atoi(studentID)
	if err = models.PutStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error in PutStudent()",
		})
		return
	}

	if err == nil {
		c.JSON(http.StatusOK, student)
		// log

	}

}
