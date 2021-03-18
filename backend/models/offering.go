package models

import "time"

type OfferingDiary struct {
	ID             int
	StudentID      int       `gorm:"not null;"`
	OfferingTypeID int       `gorm:"not null;"`
	CreatedAt      time.Time `gorm:"not null;"`
	CreatedBy      string

	Student      Student      `gorm:"references:ID"`
	OfferingType OfferingType `gorm:"references:ID"`
}

type OfferingType struct {
	ID           int
	OfferingName string `gorm:"not null;"`
}

func (OfferingDiary) TableName() string {
	return "offering_diary"
}

func (OfferingType) TableName() string {
	return "offering_type"
}
