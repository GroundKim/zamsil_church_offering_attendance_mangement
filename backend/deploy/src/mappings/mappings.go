package mappings

import (
	"zamsil_church_offering_attendance_mangement/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()
	Youth := Router.Group("Youth")

	Youth.GET("/attendances", controllers.GetAttendanceInfoByDepartment)
	Youth.POST("/attendances", controllers.PostAttendanceInfo)

	Youth.GET("/students", controllers.GetStudents)

	Youth.GET("/offering/type", controllers.GetOfferingType)
	Youth.POST("/offering", controllers.PostOffering)

	Youth.GET("/attendance/view", controllers.GetAttendanceView)
	Youth.GET("/attendance/view/excel", controllers.GetAttendanceExcel)
	Youth.GET("/offering/view", controllers.GetOfferingView)
	Youth.GET("/offering/view/excel", controllers.GetOfferingExcel)

}
