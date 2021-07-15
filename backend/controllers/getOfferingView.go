package controllers

import (
	"fmt"
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
		Month             int `json:"month"`
		WeekTotalCost     int `json:"weekTotalCost"`
		TitheTotalCost    int `json:"titheTotalCost"`
		ThanksTotalCost   int `json:"thanksTotalCost"`
		SeasonalTotalCost int `json:"seasonalTotalCost"`
		EtcTotalCost      int `json:"etcTotalCost"`
		TotalCost         int `json:"totalCost"`
	}
	year := c.Query("year")
	// get year from query
	if len(year) == 0 {
		c.JSON(400, gin.H{
			"err": "year from url paramemter is required",
		})
		return
	}

	var offeringDiariesByMonths []OfferingSummaryByMonth

	// 월 별 금액, 총 금액, 누가 가장 많이 냈는지
	for i := 1; i < 13; i++ {
		offeringDiariesByMonth := &OfferingSummaryByMonth{Month: i}

		date := year + "-" + fmt.Sprintf("%02d", i)
		parsedTime, _ := time.Parse("2006-01", date)
		var offeringDiaries []models.OfferingDiary
		models.GetOfferingDiaryByMonth(&offeringDiaries, parsedTime)

		for _, diary := range offeringDiaries {
			offeringDiariesByMonth.TotalCost += diary.Cost
			switch diary.OfferingType.OfferingTypeName {
			case "주일헌금":
				offeringDiariesByMonth.WeekTotalCost += diary.Cost
			case "십일조헌금":
				offeringDiariesByMonth.TitheTotalCost += diary.Cost
			case "감사헌금":
				offeringDiariesByMonth.ThanksTotalCost += diary.Cost
			case "절기헌금":
				offeringDiariesByMonth.SeasonalTotalCost += diary.Cost
			case "기타헌금":
				offeringDiariesByMonth.EtcTotalCost += diary.Cost
			}
		}
		offeringDiariesByMonths = append(offeringDiariesByMonths, *offeringDiariesByMonth)
	}
	c.JSON(200, offeringDiariesByMonths)
}
