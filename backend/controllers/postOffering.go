package controllers

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func PostOffering(c *gin.Context) {
	var offeringDiaries []models.OfferingDiary
	if err := c.Bind(&offeringDiaries); err != nil {
		fmt.Println("Error in json bind: post specific offering diary", err)
	}

	for _, offeringDiary := range offeringDiaries {
		offeringDiary.SaveOfferingDiary()
	}

	// log
	models.AuthToken.User.ActiveStamp("Save attendanceDiary", fmt.Sprintf("createdBy: %s", offeringDiaries[0].CreatedBy))

}
