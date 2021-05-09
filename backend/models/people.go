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

func GetTeacherByClassNameAndDepartment(teachers *[]Teacher, departmentID int, className string) (err error) {
	if err = DB.Joins("Class").Where("department_id = ?", departmentID).Find(teachers, "class.name="+"'"+className+"'").Error; err != nil {
		fmt.Println("Error in GetTeacherByClassName")
		return err
	}
	return nil
}

func GetStudentByClassNameAndDepartment(students *[]Student, departmentID int, className string) (err error) {
	if err = DB.Joins("Class").Where("department_id = ?", departmentID).Find(students, "class.name="+"'"+className+"'").Error; err != nil {
		fmt.Println("Error in GetStudentByClassName")
		return err
	}
	return nil
}

func GetTeacherByDepartment(teachers *[]Teacher, departmentID int) (err error) {
	if err = DB.Joins("Class").Find(teachers, fmt.Sprintf("department_id=%d", departmentID)).Error; err != nil {
		fmt.Println("Error in GetTeacherByDepartment")
		return err
	}
	return nil
}

func GetStudentByDepartment(students *[]Student, departmentID int) (err error) {
	if err = DB.Joins("Class").Find(students, fmt.Sprintf("department_id=%d", departmentID)).Error; err != nil {
		fmt.Println("Error in GetTeacherByDepartment")
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
