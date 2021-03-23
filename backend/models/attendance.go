package models

import "time"

type AttendanceDiary struct {
	ID        int
	StudentID int       `gorm:"not null" json:"studentId" binding:"required"`
	Date      time.Time `gorm:"not null" json:"date" binding:"required"`
	CreatedAt time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp"`
	CreatedBy string    `gorm:"not null" json:"createdBy" binding:"required"`

	Student Student `gorm:"references:ID"`
}

func (attendanceDiary *AttendanceDiary) SaveAttendanceDiary() {
	DB.Create(&attendanceDiary)
}

func (AttendanceDiary) TableName() string {
	return "attendance_diary"
}
