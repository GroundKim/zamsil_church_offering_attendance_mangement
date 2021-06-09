package models

import (
	"fmt"
	"time"
)

type Department struct {
	ID             int       `json:"deparmtentId"`
	DepartmentName int       `gorm:"not null;" json:"name"`
	CreatedAt      time.Time `gorm:"not null;" json:"createdAt"`

	//Classes []Class `gorm:"foreignKeys:DepartmentID"`
}

type Class struct {
	ID           int       `json:"classId"`
	Name         string    `gorm:"not null;" json:"name"`
	DepartmentID int       `gorm:"not null;" json:"departmentId"`
	CreatedAt    time.Time `gorm:"not null;" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"null;" json:"updatedAt"`

	Department Department `gorm:"references:ID" json:"department"`
}

func GetClasses(classes *[]Class) (err error) {
	if err = DB.Preload("Department").Find(&classes).Error; err != nil {
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
