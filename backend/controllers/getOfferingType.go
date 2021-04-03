package controllers

import (
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetOfferingType(c *gin.Context) {

	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

	var offeringTypes []models.OfferingType
	models.GetOfferingType(&offeringTypes)

	c.JSON(200, offeringTypes)
}
