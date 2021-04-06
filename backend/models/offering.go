package models

import (
	"fmt"
	"time"
)

type WeekOfferingDiary struct {
	ID        int
	Cost      int       `gorm:"not null" json:"weekOfferingCost"`
	OfferedAt time.Time `gorm:"not null" json:"offeredAt"`
	CreatedBy string    `gorm:"not null" json:"createdBy"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
}

type SpecificOfferingDiary struct {
	ID             int
	StudentID      int       `gorm:"not null;" json:"studentId"`
	OfferingTypeID int       `gorm:"not null;" json:"offeringTypeId"`
	Cost           int       `gorm:"not null;" json:"specificOfferingCost"`
	OfferedAt      time.Time `gorm:"not null" json:"offeredAt"`
	CreatedAt      time.Time `gorm:"not null;" json:"createdAt"`
	CreatedBy      string    `gorm:"not null;" json:"createdBy"`

	Student      Student      `gorm:"references:ID"`
	OfferingType OfferingType `gorm:"references:ID"`
}

type OfferingType struct {
	ID           int    `json:"offeringTypeId"`
	OfferingName string `gorm:"not null;" json:"offeringName"`
}

func (WeekOfferingDiary *WeekOfferingDiary) SaveWeekOfferingDiary() (err error) {
	if err = DB.Create(&WeekOfferingDiary).Error; err != nil {
		fmt.Println("Error in create week offering Diary")
		return err
	}
	return nil
}

func (specificOfferingDiary *SpecificOfferingDiary) SaveSpecificOfferingDiary() (err error) {

	if err = DB.Create(&specificOfferingDiary).Error; err != nil {
		fmt.Println("Error in create specific Offering Diary")
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

func (WeekOfferingDiary) TableName() string {
	return "week_offering_diary"
}

func (SpecificOfferingDiary) TableName() string {
	return "specific_offering_diary"
}

func (OfferingType) TableName() string {
	return "offering_type"
}
