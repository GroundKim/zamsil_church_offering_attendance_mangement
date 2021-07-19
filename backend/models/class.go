package models

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID             int       `json:"deparmtentId"`
	DepartmentName int       `gorm:"not null;" json:"name"`
	CreatedAt      time.Time `gorm:"not null;" json:"createdAt"`
}

type Class struct {
	ID           int       `json:"classId"`
	Name         string    `gorm:"not null;" json:"name"`
	DepartmentID int       `gorm:"not null;" json:"departmentId"`
	CreatedAt    time.Time `gorm:"not null;" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"null;" json:"updatedAt"`

	Year      int            `gorm:"not null; type:year" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Department Department `gorm:"references:ID" json:"department"`
}

func GetClasses(classes *[]Class) (err error) {
	if err = DB.Preload("Department").Find(&classes).Error; err != nil {
		fmt.Println("Error in GetClasses")
		return err
	}
	return nil
}

func GetClassesWithYear(classes *[]Class, year int) (err error) {
	if err = DB.Preload("Department").Where("year = ?", strconv.Itoa(year)).Find(&classes).Error; err != nil {
		fmt.Println("Error in GetClassesWithYear")
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
