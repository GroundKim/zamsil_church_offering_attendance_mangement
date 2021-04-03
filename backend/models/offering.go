package models

import (
	"fmt"
	"time"
)

type OfferingDiary struct {
	ID             int
	StudentID      int       `gorm:"not null;" json:"studentId"`
	OfferingTypeID int       `gorm:"not null;" json:"offeringTypeId"`
	CreatedAt      time.Time `gorm:"not null;" json:"createdAt"`
	CreatedBy      string    `gorm:"not null;" json:"createdBy"`

	Student      Student      `gorm:"references:ID"`
	OfferingType OfferingType `gorm:"references:ID"`
}

type OfferingType struct {
	ID           int    `json:"offeringTypeId"`
	OfferingName string `gorm:"not null;" json:"offeringName"`
}

func (offeringDiary *OfferingDiary) SaveOfferingDiary() (err error) {

	if err = DB.Create(&offeringDiary).Error; err != nil {
		fmt.Println("Error in create Offering Diary")
		return err
	}
	return nil
}

func GetOfferingType(OfferingTypes *[]OfferingType) (err error) {
	if err = DB.Find(OfferingTypes).Error; err != nil {
		fmt.Println("Error in GetOfferingType")
		return err
	}
	return nil
}

func (OfferingDiary) TableName() string {
	return "offering_diary"
}

func (OfferingType) TableName() string {
	return "offering_type"
}
