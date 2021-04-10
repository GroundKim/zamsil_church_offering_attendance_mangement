package models

import (
	"fmt"
	"time"
)

type OfferingDiary struct {
	ID             int
	StudentID      *int      `json:"studentId"`
	OfferingTypeID int       `gorm:"not null;" json:"offeringTypeId"`
	DepartmentID   int       `gorm:"not null;" json:"departmentId"`
	Cost           int       `gorm:"not null;" json:"offeringCost"`
	OfferedAt      time.Time `gorm:"not null" json:"offeredAt"`
	CreatedAt      time.Time `gorm:"not null;" json:"createdAt"`
	CreatedBy      string    `gorm:"not null;" json:"createdBy"`

	Student      Student      `gorm:"references:ID"`
	OfferingType OfferingType `gorm:"references:ID"`
	Department   Department   `gorm:"references:ID"`
}

type OfferingType struct {
	ID               int    `json:"offeringTypeId"`
	OfferingTypeName string `gorm:"not null;" json:"offeringName"`
}

func GetOfferingDiaryByDate(offeringDiarys *[]OfferingDiary, date time.Time) (err error) {
	theDay := date.Format("2006-01-02 ") + "00:00:00"
	theDayRange := theDay[0:11] + "23:59:59"
	if err = DB.Preload("OfferingType").Preload("Student").Where("offered_at BETWEEN ? AND ?", theDay, theDayRange).Find(&offeringDiarys).Error; err != nil {
		fmt.Println("Error in get Offering Diary by date")
		return err
	}
	return nil
}

func (OfferingDiary *OfferingDiary) SaveOfferingDiary() (err error) {
	if err = DB.Create(&OfferingDiary).Error; err != nil {
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
