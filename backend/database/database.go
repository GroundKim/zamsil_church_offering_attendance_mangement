package database

import (
	"fmt"
	"zamsil_church_offering_attendance_mangement/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb(conf *models.Config) *gorm.DB {
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

	return db
}
