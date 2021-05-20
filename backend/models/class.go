package models

import (
	"fmt"
	"time"
)

type Department struct {
	ID             int
	DepartmentName int       `gorm:"not null;"`
	CreatedAt      time.Time `gorm:"not null;"`

	Classes []Class `gorm:"foreignKeys:DepartmentID"`
}

type Class struct {
	ID           int
	Name         string    `gorm:"not null;"`
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

func GetClassesByDepartment(classes *[]Class, departmentID int) (err error) {
	if err = DB.Where("department_id = ?", departmentID).Find(classes).Error; err != nil {
		fmt.Println("Error in GetClassesByDepartment")
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
