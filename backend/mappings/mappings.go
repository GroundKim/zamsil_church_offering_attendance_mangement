package mappings

import (
	"zamsil_church_offering_attendance_mangement/backend/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()

	Youth := Router.Group("Youth")
	Youth.GET("/input/1", controllers.GetAttendanceInfo)

}
