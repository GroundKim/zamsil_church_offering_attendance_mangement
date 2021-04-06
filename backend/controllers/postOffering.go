package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func PostOffering(c *gin.Context) {

	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	offeringType := c.Query("offering_type")
	if offeringType == "specific" {
		var specificOfferingDiarys []models.SpecificOfferingDiary
		if err := c.Bind(&specificOfferingDiarys); err != nil {
			fmt.Println("Error in json bind: post specific offering diary", err)
		}

		for _, offeringDiary := range specificOfferingDiarys {
			offeringDiary.SaveSpecificOfferingDiary()
		}
		return
	}

	if offeringType == "week" {
		var WeekOfferingDiary models.WeekOfferingDiary
		if err := c.Bind(&WeekOfferingDiary); err != nil {
			fmt.Println("Error in json bind: post week offering diary", err)
		}
		WeekOfferingDiary.SaveWeekOfferingDiary()
		return
	}

}
