package mappings

import (
	"zamsil_church_offering_attendance_mangement/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Youth := Router.Group("Youth")
	Youth.GET("/resource/attendance/department/1", controllers.GetAttendanceInfoByDepartment(1))
	Youth.GET("/resource/attendance/department/2", controllers.GetAttendanceInfoByDepartment(2))
}
