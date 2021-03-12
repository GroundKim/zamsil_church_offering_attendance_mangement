package main

import (
	"zamsil_church_offering_attendance_mangement/backend/config"
	"zamsil_church_offering_attendance_mangement/backend/database"
)

// Main program.
func main() {
	conf := config.GetConf("dev")
	db := database.InitDb(conf)

}
