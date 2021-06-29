package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

// save attendance information into database
func PostAttendanceDiary(c *gin.Context) {
	var attendanceDiaries []models.AttendanceDiary
	if err := c.Bind(&attendanceDiaries); err != nil {
		fmt.Println("Error in json bind: post attendance diary ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "fail in binding attendanceDiary",
		})
		return
	}

	for _, attendanceDiary := range attendanceDiaries {
		// handling post attendanceDiary from attendanceViewDetail.vue
		// and, if createdBy is null
		if attendanceDiary.CreatedBy == "" {
			attendanceDiary.CreatedBy = models.AuthToken.User.Name
		}

		if err := attendanceDiary.SaveAttendanceDiary(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "fail in save attendance Diary",
			})
			return
		}
	}

	// log
	models.AuthToken.User.ActiveStamp("Save attendanceDiary", "")
}

func DeleteAttendanceDiary(c *gin.Context) {
	var attendanceDiary models.AttendanceDiary
	var err error
	attendanceDiary.StudentID, err = strconv.Atoi(c.Query("student_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "studnet_id is required",
		})
		return
	}

	attendanceDiary.AttendedAt, err = time.Parse("2006-01-02", c.Param("date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "date is required",
		})
		return
	}

	if err = attendanceDiary.DeleteAttendanceDiaryByStudentIDAndAttendedAt(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": "fail in deleteAttendanceDiary",
		})
		return

	}
	c.JSON(http.StatusOK, attendanceDiary)
}
