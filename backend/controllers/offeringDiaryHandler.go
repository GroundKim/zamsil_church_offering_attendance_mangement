package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func ChangeOfferingDiaryTypeAndCost(c *gin.Context) {
	var offeringDiary models.OfferingDiary
	if err := c.BindJSON(&offeringDiary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "fail in binding offeringDiary",
		})
		fmt.Println(err)
		return
	}

	if err := models.ChangeOfferingDiaryTypeAndCost(&offeringDiary); err != nil {
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

func GetOfferingDiarySummary(c *gin.Context) {
	date := c.Param("date")
	parsedDate, _ := time.Parse("2006-01-02", date)
	offeringCostWrappers, err := models.GetSummaryOfferingCostByDate(parsedDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "fail in get summary offering Cost By date",
		})
		return
	}

	c.JSON(http.StatusOK, offeringCostWrappers)
}

func PostOfferingDiary(c *gin.Context) {
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

func GetOfferingDiaryType(c *gin.Context) {
	var offeringTypes []models.OfferingType
	models.GetOfferingType(&offeringTypes)

	c.JSON(200, offeringTypes)
}
