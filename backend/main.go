package main

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"
)

// Main program.
func main() {
	conf := models.GetConf("dev")

	// db := database.InitDb(conf)
	fmt.Println(conf)

}
