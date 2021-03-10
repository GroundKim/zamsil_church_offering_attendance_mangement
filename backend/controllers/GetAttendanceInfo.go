package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetAttendanceInfo(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"count":        15,
		"teacherName":  "선생님 예시",
		"studentNames": "김지상",
	})
}
