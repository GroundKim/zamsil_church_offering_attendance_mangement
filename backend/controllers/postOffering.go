package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func PostOffering(c *gin.Context) {

	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	var offeringDiarys []models.OfferingDiary
	if err := c.Bind(&offeringDiarys); err != nil {
		fmt.Println("Error in json bind: post offering diary", err)
	}

	for _, offeringDiary := range offeringDiarys {
		offeringDiary.SaveOfferingDiary()
	}

}
