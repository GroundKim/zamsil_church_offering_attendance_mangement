package main

import (
	"zamsil_church_offering_attendance_mangement/config"
	"zamsil_church_offering_attendance_mangement/mappings"
	"zamsil_church_offering_attendance_mangement/models"
)

// Main program.
func main() {
	conf := config.GetConf("dev")
	models.InitDb(conf)

	mappings.CreateUrlMappings()
	mappings.Router.Run()
}
