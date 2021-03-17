package models

import "time"

type Student struct {
	ID           int
	Name         string    `gorm:"not null;"`
	DepartmentId int       `gorm:"not null;"`
	ClassID      int       `gorm:"not null;"`
	CreatedAt    time.Time `gorm:"not null;"`

	Class      Class      `gorm:"references:ID"`
	Department Department `gorm:"references:ID"`
}

type Teacher struct {
	ID           int
	Name         string    `gorm:"not null;"`
	DepartmentID int       `gorm:"not null;"`
	ClassId      int       `gorm:"not null;"`
	CreatedAt    time.Time `gorm:"not null;"`

	Class      Class      `gorm:"references:ID"`
	Department Department `gorm:"references:ID"`
}

func (Student) TableName() string {
	return "student"
}

func (Teacher) TableName() string {
	return "teacher"
}
