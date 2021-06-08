package controllers

import (
	"net/http"
	"strconv"
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetAbsenceDiaries(c *gin.Context) {
	date := c.Query("date")
	var absenceDiaries []models.AbsenceDiary

	if len(date) > 0 {
		parsedDate, _ := time.Parse("2006-01-02", date)
		models.GetAbsenceDiariesByDate(&absenceDiaries, parsedDate)
	}
	c.JSON(200, absenceDiaries)
}

func GetAbsenceType(c *gin.Context) {
	var absenceTypes []models.AbsenceType

	if err := models.GetAbsenceTypes(&absenceTypes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "fail in getAbsenceTypes",
		})
	}

	c.JSON(200, absenceTypes)
}

func PostAbsentStudents(c *gin.Context) {
	var absentStudents []models.AbsenceDiary

	if err := c.BindJSON(&absentStudents); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"err": "binding error",
		})
		return
	}

	for _, absentStudent := range absentStudents {
		if err := absentStudent.SaveAbsentDiary(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "fail in saveAbsenceDiary",
			})
			return
		}
	}

	c.JSON(200, absentStudents)
}

func DeleteAbsenceDiary(c *gin.Context) {
	absenceDiaryID := c.Param("id")
	parsedID, _ := strconv.Atoi(absenceDiaryID)
	absenceDiary := &models.AbsenceDiary{ID: &parsedID}

	if err := absenceDiary.DeleteAbsenceDiaryByID(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "fail in DeleteAbsenceDiary",
		})
		return
	}
	c.JSON(http.StatusOK, absenceDiary)
}
