package controllers

import (
	"net/http"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
	var classes []models.Class
	if err := models.GetClasses(&classes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error in GetClass()",
		})
	}

	c.JSON(http.StatusOK, classes)
}

func PostClasses(c *gin.Context) {
	var newClass models.Class
	var classes []models.Class
	hasSameClass := false

	if err := models.GetClasses(&classes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error in GetClass()",
		})
	}

	for _, class := range classes {
		if class.Name == newClass.Name {
			hasSameClass = true
		}
	}

	if hasSameClass {
		// patch
	} else {
		// save new class
	}
	// 지석이한테 2021학년도 라벨링 해야 되는지 물어보기
	// 문법 이상 있는지 없는지 물어보기
}
