package controllers

import (
	"fmt"
	"time"
	"zamsil_church_offering_attendance_mangement/controllers/makeExcel"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetAttendanceExcel(c *gin.Context) {
	date := c.Query("date")
	parsedDate, _ := time.Parse("2006-01-02", date)
	if len(date) != 0 {
		makeExcel.SaveAttendaceViewExcel(parsedDate)
		c.Writer.Header().Add("Content-type", "application/octet-stream")
		date := c.Query("date")
		c.FileAttachment(fmt.Sprintf("./data/attendanceDiary/excel/%s.xlsx", date), fmt.Sprintf("출석부_%s.xlsx", date))
		models.AuthToken.User.ActiveStamp("Download", fmt.Sprintf("./data/attendanceDiary/excel/%s.xlsx", date))
		return
	}
	c.AbortWithStatus(500)
}
