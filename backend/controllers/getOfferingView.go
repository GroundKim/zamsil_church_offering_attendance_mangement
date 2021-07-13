package controllers

import (
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetOfferingViewList(c *gin.Context) {
	if len(c.Query("year")) != 0 {
		year := c.Query("year")
		parsedYear, _ := time.Parse("2006", year)
		var offeringDiaries []models.OfferingDiary
		var offeringAts []time.Time
		models.GetOfferingDiaryByYear(&offeringDiaries, parsedYear)

		for i := 0; i < len(offeringDiaries); i++ {
			isFound := false
			for j := 0; j < len(offeringAts); j++ {
				if offeringDiaries[i].OfferedAt.Format("2006-01-02") == offeringAts[j].Format("2006-01-02") {
					isFound = true
				}
			}

			if !isFound {
				offeringAts = append(offeringAts, offeringDiaries[i].OfferedAt)
			}

		}
		payload := map[string][]time.Time{"offeredAts": offeringAts}
		c.JSON(200, payload)
	}
}

func GetOfferingView(c *gin.Context) {
	date := c.Query("date")
	parsedDate, _ := time.Parse("2006-01-02", date)
	var offeringDiaries []models.OfferingDiary

	models.GetOfferingDiaryByDate(&offeringDiaries, parsedDate)
	c.JSON(200, offeringDiaries)
}

func GetOfferingViewStatisticSummaryByYear(c *gin.Context) {
	type OfferingSummaryByMonth struct {
		month int
		weekTotalCost int
		titheTotalCost int
		seasonalTotalCost int
		etcTotalCost int
	}
	year := c.Query("year")
	// get year from query
	if len(year) == 0 {
		c.JSON(400, gin.H{
			"err": "year from url paramemter is required",
		})
		return
	}

	// 월 별 금액, 총 금액, 누가 가장 많이 냈는지
	var offeringDiaries []models.OfferingDiary
	parsedTime, _ := time.Parse("2006", year)
	models.GetOfferingDiaryByYear(&offeringDiaries, parsedTime)

	for _, diary := range offeringDiaries {
		diary.
	}

	c.JSON(200, offeringDiaries)
}
