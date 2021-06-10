package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func PutOfferingDiary(c *gin.Context) {
	var offeringDiary models.OfferingDiary
	if err := c.BindJSON(&offeringDiary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "fail in binding offeringDiary",
		})
		fmt.Println(err)
		return
	}

	if err := models.PutOfferingDiary(&offeringDiary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "fail in put offering diary",
		})
		return
	}
	c.JSON(http.StatusOK, offeringDiary)
}

func DeleteOfferingDiary(c *gin.Context) {
	offeringDiaryID := c.Param("id")
	parsedOfferingDiaryID, _ := strconv.Atoi(offeringDiaryID)
	offeringDiary := &models.OfferingDiary{ID: parsedOfferingDiaryID}
	if err := models.DeleteOfferingDiary(offeringDiary); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "fail in delete offering diary",
		})
		return
	}

	c.JSON(http.StatusOK, offeringDiary)
}
