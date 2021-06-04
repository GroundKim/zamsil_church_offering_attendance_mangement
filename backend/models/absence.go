package models

import "time"

type AbsenceDiary struct {
	ID            int       `json:"absenceDiaryId"`
	StudentID     int       `gorm:"not null" json:"studentId"`
	AbsentAt      time.Time `gorm:"not null; type:date" json:"absentAt"`
	AbsenceTypeID int       `gorm:"not null;" json:"absenceTypeId"`
	Reason        string    `json:"reason"`
	CreatedAt     time.Time `gorm:"not null" sql:"DEFAULT:current_timestamp" json:"createdAt"`
	CreatedBy     string    `gorm:"not null" json:"createdBy"`

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

func (AbsenceDiary) TableName() string {
	return "absence_diary"
}
