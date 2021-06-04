package controllers

import (
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
