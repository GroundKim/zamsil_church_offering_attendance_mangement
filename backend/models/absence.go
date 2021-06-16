package models

import (
	"time"
)

type AbsenceDiary struct {
	ID            *int      `json:"absenceDiaryId"`
	StudentID     int       `gorm:"not null" json:"studentId"`
	AbsentAt      time.Time `gorm:"not null; type:date" json:"absentAt"`
	AbsenceTypeID int       `gorm:"not null;" json:"absenceTypeId"`
	Reason        string    `json:"absenceReason"`
	CreatedAt     time.Time `gorm:"not null; default:current_timestamp(3)" json:"createdAt"`
	CreatedBy     string    `gorm:"null" json:"createdBy"`

	Student     Student     `gorm:"references:ID" json:"-"`
	AbsenceType AbsenceType `gorm:"references:ID" json:"absenceType"`
}

type AbsenceType struct {
	ID   int    `json:"absenceTypeId"`
	Name string `gorm:"not null" json:"name"`
}

func GetAbsenceDiariesByDate(absences *[]AbsenceDiary, date time.Time) (err error) {
	if err = DB.Preload("AbsenceType").Where("absent_at = ?", date.Format("2006-01-02")).Find(absences).Error; err != nil {
		return err
	}
	return nil
}

func GetAbsenceTypes(absenceTypes *[]AbsenceType) (err error) {
	if err = DB.Find(&absenceTypes).Error; err != nil {
		return err
	}
	return nil
}

func (absenceDiary *AbsenceDiary) SaveAbsentDiary() (err error) {
	if err = DB.Create(&absenceDiary).Error; err != nil {
		return err
	}
	return nil
}

func (absenceDiary *AbsenceDiary) DeleteAbsenceDiaryByID() (err error) {
	if err = DB.Delete(&absenceDiary).Error; err != nil {
		return err
	}
	return nil
}

func (absenceDiary *AbsenceDiary) UpdateAbsenceDiaryByID() (err error) {
	if err = DB.Save(&absenceDiary).Error; err != nil {
		return err
	}
	return nil
}

func (absenceDiary *AbsenceDiary) UpdateAbsenceDiaryReason() (err error) {
	if err = DB.Model(&absenceDiary).Update("reason", absenceDiary.Reason).Error; err != nil {
		return err
	}
	return nil
}

func (AbsenceDiary) TableName() string {
	return "absence_diary"
}

func (AbsenceType) TableName() string {
	return "absence_type"
}
