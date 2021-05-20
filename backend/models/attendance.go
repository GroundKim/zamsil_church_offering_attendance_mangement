package models

import (
	"fmt"
	"time"
)

type AttendanceDiary struct {
	ID         int
	StudentID  int       `gorm:"not null" json:"studentId" binding:"required"`
	AttendedAt time.Time `gorm:"not null; type:date" json:"attendedAt" binding:"required"`
	CreatedAt  time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp"`
	CreatedBy  string    `gorm:"not null" json:"createdBy" binding:"required"`

	Student Student `gorm:"references:ID"`
}

func (attendanceDiary *AttendanceDiary) SaveAttendanceDiary() {
	DB.Create(&attendanceDiary)
}

func GetAttendanceViewByDate(AttendanceDiaries *[]AttendanceDiary, date time.Time) (err error) {
	theYear := date.Format("2006-01-02 ") + "00:00:00"
	theYearRange := theYear[0:10] + " 23:59:59"
	if err = DB.Preload("Student").Preload("Student.Class").Where("attended_at BETWEEN ? AND ?", theYear, theYearRange).Find(&AttendanceDiaries).Error; err != nil {
		fmt.Println("Error in get attendance view by date")
		return err
	}
	return nil
}

func GetAttendanceViewByYear(AttendanceDiaries *[]AttendanceDiary, date time.Time) (err error) {
	theYear := date.Format("2006-01-02 ") + "00:00:00"
	theYearRange := theYear[0:4] + "-12-31 23:59:59"
	if err = DB.Preload("Student").Preload("Student.Class").Where("attended_at BETWEEN ? AND ?", theYear, theYearRange).Find(&AttendanceDiaries).Error; err != nil {
		fmt.Println("Error in get attendance view by date")
		return err
	}
	return nil
}

//... use interface and reciever function

// func (AttendanceDiarys *AttendanceDiary) GetAttendanceViewByYear(date time.Time) (err error) {
// 	theYear := date.Format("2006-01-02 ") + "00:00:00"
// 	theYearRange := theYear[0:4] + "-12-31 11:59:59"
// 	if err = DB.Preload("Student").Where("created_at BETWEEN ? AND ?", theYear, theYearRange).Find(&AttendanceDiarys).Error; err != nil {
// 		fmt.Println("Error in get attendance view by date")
// 		return err
// 	}
// 	return nil
// }

func (AttendanceDiary) TableName() string {
	return "attendance_diary"
}
