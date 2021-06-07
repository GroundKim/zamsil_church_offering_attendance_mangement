package controllers

import (
	"net/http"
	"strconv"
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/gin-gonic/gin"
)

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
