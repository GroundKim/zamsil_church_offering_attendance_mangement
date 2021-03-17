package models

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb(conf *config.Config) {
	dbCredentials := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dbCredentials))
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to the database!")
	}
	Migrate(db)
	DB = db
}

// Migrate models using GORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		OfferingDiary{},
		OfferingType{},
		Student{},
		Teacher{},
		AttendanceDiary{},
		Class{},
		Department{},
	)
	db.Commit()
	fmt.Println("Auto Migration has been excuted")
}

func GetDB() *gorm.DB {
	return DB
}
