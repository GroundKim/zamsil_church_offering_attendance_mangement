package models

import (
	"fmt"
	"time"
)

type AttendanceDiary struct {
	ID         int       `json:"attendanceDiaryId"`
	StudentID  int       `gorm:"not null" json:"studentId" binding:"required"`
	AttendedAt time.Time `gorm:"not null; type:date" json:"attendedAt" binding:"required"`
	CreatedAt  time.Time `gorm:"not null default: current_timestamp(3)"`
	CreatedBy  string    `gorm:"not null" json:"createdBy"`

	Student Student `gorm:"references:ID" json:"student"`
}

func (attendanceDiary *AttendanceDiary) SaveAttendanceDiary() (err error) {
	if err = DB.Create(&attendanceDiary).Error; err != nil {
		return err
	}
	return nil
}

func (attendanceDiary *AttendanceDiary) DeleteAttendanceDiaryByStudentIDAndAttendedAt() (err error) {
	if err = DB.Exec(fmt.Sprintf("DELETE FROM attendance_diary WHERE student_id = %d && attended_at = '%s'", attendanceDiary.StudentID, attendanceDiary.AttendedAt.Format("2006-01-02"))).Error; err != nil {
		return err
	}
	return nil
}

func GetAttendanceDiariesByDate(attendanceDiaries *[]AttendanceDiary, date time.Time) (err error) {
	parsedDate := date.Format("2006-01-02 ") + "00:00:00"
	parsedDateRange := parsedDate[0:10] + " 23:59:59"
	if err = DB.Preload("Student").Where("attended_at BETWEEN ? AND ?", parsedDate, parsedDateRange).Find(&attendanceDiaries).Error; err != nil {
		fmt.Println("Error in get attendance view by date")
		return err
	}
	return nil
}

func GetAttendanceDiariesByMonth(attendanceDiaries *[]AttendanceDiary, month time.Time) (err error) {
	parsedMonth := month.Format("2006-01") + "-01 00:00:00"
	searchMonth := int(month.Month())

	parsedMonthRange := parsedMonth[0:5] + fmt.Sprintf("%02d", searchMonth) + "-31 23:59:59"

	if err = DB.Preload("Student").Preload("Student.Class").Where("attended_at BETWEEN ? AND ?", parsedMonth, parsedMonthRange).Find(&attendanceDiaries).Error; err != nil {
		fmt.Println("Error in get attendance view by month")
		return err
	}
	return nil
}

func GetAttendanceViewByDateWithoutDuplicatedStudentID(AttendanceDiaries *[]AttendanceDiary, date time.Time) (err error) {
	theYear := date.Format("2006-01-02 ") + "00:00:00"
	theYearRange := theYear[0:10] + " 23:59:59"
	if err = DB.Preload("Student").Preload("Student.Class").Where("attended_at BETWEEN ? AND ?", theYear, theYearRange).Group("student_id").Find(&AttendanceDiaries).Error; err != nil {
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
