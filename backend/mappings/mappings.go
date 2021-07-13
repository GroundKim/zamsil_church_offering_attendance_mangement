package mappings

import (
	"io"
	"os"
	"time"
	"zamsil_church_offering_attendance_mangement/config"
	"zamsil_church_offering_attendance_mangement/controllers"
	"zamsil_church_offering_attendance_mangement/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings(conf *config.Config) {
	// gin logging
	now := time.Now().Format("2006-01-02")
	f, _ := os.Create("data/log/" + now + "_gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	Router = gin.Default()

	// middelware for cors
	Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{conf.CORS.ALLOWORIGINS},
		AllowCredentials: conf.CORS.ACCESSCONTROLALLOWCREDENTIALS,
		AllowHeaders:     []string{"content-type"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE"},
	}))

	// Router without authorization process
	Router.POST("/Youth/login", controllers.Login(conf))
	Router.GET("/Youth/login", controllers.ValidateUser(conf))

	// Router Get with middlewares AUthorize
	Youth := Router.Group("Youth").Use(middlewares.Authorize(conf))

	Youth.GET("/attendances", controllers.GetAttendanceInfoByDepartment)
	Youth.POST("/attendances", controllers.PostAttendanceDiary)
	Youth.DELETE("/attendance/:date", controllers.DeleteAttendanceDiary)

	Youth.GET("/absence", controllers.GetAbsenceDiaries)
	Youth.POST("/absence", controllers.PostAbsentStudents)
	Youth.GET("/absence/types", controllers.GetAbsenceType)
	Youth.DELETE("/absence/:id", controllers.DeleteAbsenceDiary)
	Youth.PUT("/absence/:id", controllers.PutAbsenceDiary)
	Youth.PATCH("/absence/:id", controllers.PatchAbsenceDiary)

	Youth.GET("/attendance/view", controllers.GetAttendanceView)
	Youth.GET("/attendance/view/list", controllers.GetAttendanceViewList)
	Youth.GET("/attendance/view/excel", controllers.GetAttendanceExcel)
	Youth.GET("/attendance/view/statistic/number/month", controllers.GetAttendanceNumberByMonth)

	Youth.GET("/members", controllers.GetMembers)

	Youth.GET("/students", controllers.GetStudents)
	Youth.POST("/students", controllers.SaveStudents)
	Youth.PUT("/students/:id", controllers.PutStudent)
	Youth.DELETE("/students/:id", controllers.DeleteStudent)

	Youth.GET("/teachers", controllers.GetTeachers)
	Youth.POST("/teachers", controllers.SaveTeachers)

	Youth.POST("/offering", controllers.PostOfferingDiary)
	Youth.GET("/offering/types", controllers.GetOfferingDiaryType)
	Youth.GET("/offering/view", controllers.GetOfferingView)
	Youth.GET("/offering/view/list", controllers.GetOfferingViewList)
	Youth.GET("/offering/view/excel", controllers.GetOfferingExcel)
	Youth.GET("/offering/summary/:date", controllers.GetOfferingDiarySummary)
	Youth.PUT("/offering", controllers.ChangeOfferingDiaryTypeAndCost)
	Youth.DELETE("/offering/:id", controllers.DeleteOfferingDiary)
	Youth.GET("/offering/view/statistic/summary", controllers.GetOfferingViewStatisticSummaryByYear)

	Youth.GET("/classes", controllers.GetClasses)

}
