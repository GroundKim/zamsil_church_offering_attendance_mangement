package models

import (
	"fmt"
	"time"
)

type Department struct {
	ID           int
	DepartmentID int       `gorm:"not null;"`
	CreatedAt    time.Time `gorm:"not null;"`

	Classes []Class `gorm:"foreignKeys:DepartmentID"`
}

type Class struct {
	ID           int
	ClassName      string    `gorm:"not null;"`
	DepartmentID int       `gorm:"not null;"`
	CreatedAt    time.Time `gorm:"not null;"`
}

func GetClasses(classes *[]Class) (err error) {
	if err = DB.Find(classes).Error; err != nil {
		fmt.Println("Error in GetClasses")
		return err
	}
	return nil
}

func (Department) TableName() string {
	return "department"
}

func (Class) TableName() string {
	return "class"
}
