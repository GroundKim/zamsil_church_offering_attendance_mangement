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
		conf.DATABASE.DB_USERNAME,
		conf.DATABASE.DB_PASSWORD,
		conf.DATABASE.DB_HOST,
		conf.DATABASE.DB_PORT,
		conf.DATABASE.DB_NAME,
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
		Department{},
		Class{},
		Student{},
		Teacher{},
		OfferingType{},
		OfferingDiary{},
		AttendanceDiary{},
		User{},
		authToken{},
		AbsenceDiary{},
		Log{},
	)
	fmt.Println("Auto Migration has been executed")
}

func GetDB() *gorm.DB {
	return DB
}
