package models

import (
	"fmt"
	"time"
)

type Student struct {
	ID        int       `json:"studentId"`
	Name      string    `gorm:"not null;" json:"name"`
	ClassID   int       `gorm:"not null;" json:"className"`
	CreatedAt time.Time `gorm:"not null;" json:"createdAt" `

	Class Class `gorm:"references:ID" json:"-"`
}

type Teacher struct {
	ID        int
	Name      string    `gorm:"not null;"`
	ClassID   int       `gorm:"not null;"`
	CreatedAt time.Time `gorm:"not null;"`

	Class Class `gorm:"references:ID"`
}

func GetStudents(students *[]Student) (err error) {
	if err = DB.Find(students).Error; err != nil {
		fmt.Println("Error in GetStudent")
		return err
	}
	return nil
}

func GetTeachers(teachers *[]Teacher) (err error) {
	if err = DB.Find(teachers).Error; err != nil {
		fmt.Println("Error in getAllTeacher")
		return err
	}
	return nil
}

func GetTeacherByDeaprtment(teachers *[]Teacher, departmentID int) (err error) {
	if err = DB.Joins("Class").Find(teachers, fmt.Sprintf("class.department_id=%d", departmentID)).Error; err != nil {
		fmt.Println("Error in GetTeacherByDeaprtment")
		return err
	}
	return nil
}

func GetStudentByDepartment(students *[]Student, departmentID int) (err error) {
	if err = DB.Joins("Class").Find(students, fmt.Sprintf("class.department_id=%d", departmentID)).Error; err != nil {
		fmt.Println("Error in GetTeacherByDeaprtment")
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
