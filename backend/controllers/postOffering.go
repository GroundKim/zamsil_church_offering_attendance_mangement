package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func PostOffering(c *gin.Context) {
	var offeringDiarys []models.OfferingDiary
	if err := c.Bind(&offeringDiarys); err != nil {
		fmt.Println("Error in json bind: post specific offering diary", err)
	}

	for _, offeringDiary := range offeringDiarys {
		offeringDiary.SaveOfferingDiary()
	}

}
