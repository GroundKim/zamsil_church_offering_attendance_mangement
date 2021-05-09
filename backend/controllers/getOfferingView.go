package controllers

import (
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

func GetOfferingView(c *gin.Context) {

	if len(c.Query("year")) != 0 {

		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")

		year := c.Query("year")
		parsedYear, _ := time.Parse("2006", year)
		var offeringDiaries []models.OfferingDiary
		var offeringAts []time.Time
		models.GetOfferingDiaryByYear(&offeringDiaries, parsedYear)

		for i := 0; i < len(offeringDiaries); i++ {
			isFound := false
			for j := 0; j < len(offeringDiaries); j++ {

				if offeringDiaries[i].CreatedAt.Format("2006-01-02") == offeringAts[j].Format("2006-01-02") {
					isFound = true
				}
			}

			if !isFound {
				offeringAts = append(offeringAts, offeringDiaries[i].CreatedAt)
			}

		}
		payload := map[string][]time.Time{"offeredAts": offeringAts}
		c.JSON(200, payload)
	}

}