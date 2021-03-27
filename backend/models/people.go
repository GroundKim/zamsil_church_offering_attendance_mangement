package models

import (
	"fmt"
	"time"
)

type Student struct {
	ID           int       `json:"studentId"`
	Name         string    `gorm:"not null;" json:"name"`
	Grade        int       `gorm:"not null;" json:"grade"`
	DepartmentId int       `gorm:"not null;" json:"departmentId"`
	ClassID      int       `gorm:"not null;" json:"classId"`
	CreatedAt    time.Time `gorm:"not null;" json:"createdAt" `

	Class      Class      `gorm:"references:ID" json:"-"`
	Department Department `gorm:"references:ID" json:"-"`
}

type Teacher struct {
	ID           int
	Name         string    `gorm:"not null;"`
	DepartmentID int       `gorm:"not null;"`
	ClassID      int       `gorm:"not null;"`
	CreatedAt    time.Time `gorm:"not null;"`

	Class      Class      `gorm:"references:ID"`
	Department Department `gorm:"references:ID"`
}

func GetAllStudent(students *[]Student) (err error) {
	if err = DB.Find(students).Error; err != nil {
		fmt.Println("Error in GetAllStudent")
		return err
	}
	return nil
}

func GetAllTeacher(teachers *[]Teacher) (err error) {
	if err = DB.Find(teachers).Error; err != nil {
		fmt.Println("Error in getAllTeacher")
		return err
	}
	return nil
}

func GetTeacherByDeaprtment(teachers *[]Teacher, departmentID int) (err error) {
	if err = DB.Where("department_id = ?", departmentID).Find(teachers).Error; err != nil {
		fmt.Println("Error in getAllTeacher")
		return err
	}
	return nil
}

func GetStudentByDepartment(students *[]Student, departmentID int) (err error) {
	if err = DB.Where("department_id = ?", departmentID).Find(students).Error; err != nil {
		fmt.Println("Error in GetStudentByDepartment")
		return err
	}
	return nil
}

func (Student) TableName() string {
	return "student"
}

func (Teacher) TableName() string {
	return "teacher"
}
