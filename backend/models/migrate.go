package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrate models using GORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&OfferingDiary{},
		&OfferingType{},
		&Student{},
		&Teacher{},
		&AttendanceDiary{},
		&Class{},
		&Department{},
	)
	db.Commit()
	fmt.Println("Auto Migration has been excuted")
}
