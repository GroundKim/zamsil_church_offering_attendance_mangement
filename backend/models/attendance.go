package models

import "time"

type AttendanceDiary struct {
	ID        int
	StudentID int       `gorm:"not null;"`
	Date      time.Time `gorm:"not null;"`
	CreatedAt time.Time `gorm:"not null;"`
	CreatedBy string    `gorm:"not null;"`

	Student Student `gorm:"references:ID"`
}

func (AttendanceDiary) TableName() string {
	return "attendance_diary"
}
