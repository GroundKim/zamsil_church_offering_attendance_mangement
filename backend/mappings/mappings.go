package mappings

import (
	"zamsil_church_offering_attendance_mangement/config"
	"zamsil_church_offering_attendance_mangement/controllers"
	"zamsil_church_offering_attendance_mangement/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings(conf *config.Config) {
	Router = gin.Default()
	// middelware for cors
	Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{conf.CORS.ALLOWORIGINS},
		AllowCredentials: conf.CORS.ACCESSCONTROLALLOWCREDENTIALS,
		AllowHeaders:     []string{"content-type"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE"},
	}))

	Router.POST("/Youth/login", controllers.Login(conf))

	// Router Get with middlewares AUthorize
	Youth := Router.Group("Youth").Use(middlewares.Authorize(conf))

	Youth.GET("/attendances", controllers.GetAttendanceInfoByDepartment)
	Youth.POST("/attendances", controllers.PostAttendanceInfo)

	Youth.GET("/absence", controllers.GetAbsenceDiaries)

	Youth.GET("/attendance/view/list", controllers.GetAttendanceViewList)
	Youth.GET("/attendance/view/excel", controllers.GetAttendanceExcel)
	Youth.GET("/attendance/view", controllers.GetAttendanceView)

	Youth.GET("/members", controllers.GetMembers)

	Youth.GET("/students", controllers.GetStudents)
	Youth.POST("/students", controllers.SaveStudents)
	Youth.PUT("/students/:id", controllers.PutStudent)
	Youth.DELETE("/students/:id", controllers.DeleteStudent)

	Youth.POST("/teachers", controllers.SaveTeachers)

	Youth.GET("/offering/type", controllers.GetOfferingType)
	Youth.POST("/offering", controllers.PostOffering)

	Youth.GET("/offering/view", controllers.GetOfferingView)
	Youth.GET("/offering/view/excel", controllers.GetOfferingExcel)

	Youth.GET("/classes", controllers.GetClasses)

}
