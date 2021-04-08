package controllers

import (
	"fmt"
	"time"
	"zamsil_church_offering_attendance_mangement/controllers/makeExcel"

	"github.com/gin-gonic/gin"
)

func GetOfferingExcel(c *gin.Context) {

	makeExcel.SaveExcel(time.Now(), "작성자입니다") //...  get createdBy from gorm(mysql)

	if len(c.Query("date")) != 0 {

		c.Writer.Header().Add("Content-type", "application/octet-stream")

		date := c.Query("date")
		c.FileAttachment(fmt.Sprintf("./data/offeringDiary/excel/%s.xlsx", date), fmt.Sprintf("헌금통계표 %s.xlsx", date))
		return
	}

	c.AbortWithStatus(500)
}
