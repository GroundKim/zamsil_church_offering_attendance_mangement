package controllers

import (
	"fmt"
	"time"
	"zamsil_church_offering_attendance_mangement/controllers/makeExcel"

	"github.com/gin-gonic/gin"
)

func GetOfferingExcel(c *gin.Context) {
	date := c.Query("date")
	parsedDate, _ := time.Parse("2006-01-02", date)
	if len(date) != 0 {
		fmt.Println(parsedDate)
		makeExcel.SaveOfferingViewExcel(parsedDate, "작성자입니다") //...  get createdBy from gorm(mysql)
		c.Writer.Header().Add("Content-type", "application/octet-stream")
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		date := c.Query("date")
		c.FileAttachment(fmt.Sprintf("./data/offeringDiary/excel/%s.xlsx", date), fmt.Sprintf("헌금통계표_%s.xlsx", date))
		return
	}

	c.AbortWithStatus(500)
}
