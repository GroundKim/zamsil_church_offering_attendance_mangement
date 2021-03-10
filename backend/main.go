package main

import "zamsil_church_offering_attendance_mangement/backend/mappings"

func main() {
	mappings.CreateUrlMappings()
	mappings.Router.Run()

}
