package models

// member

import (
	"fmt"
	"time"
)

type Student struct {
	ID                int       `json:"studentId"`
	Name              string    `gorm:"not null;" json:"name"`
	ClassID           int       `gorm:"not null;" json:"classId"`
	CreatedAt         time.Time `gorm:"not null;" json:"createdAt"`
	DayOfBirth        time.Time `gorm:"type:date;" json:"dayOfBirth"`
	PhoneNumber       string    `json:"phoneNumber"`
	ParentPhoneNumber string    `json:"parentPhoneNumber"`
	SchoolName        string    `json:"schoolName"`

	Class Class `gorm:"references:ID" json:"-"`
}

type Teacher struct {
	ID        int
	Name      string    `gorm:"not null;"`
	ClassID   int       `gorm:"not null;"`
	CreatedAt time.Time `gorm:"not null;"`

	Class Class `gorm:"references:ID"`
}

type StudentsWithDepartment struct {
	ID             int       `json:"studentId"`
	Name           string    `json:"name"`
	ClassID        int       `json:"classId"`
	CreatedAt      time.Time `json:"createdAt"`
	DepartmentName int       `json:"departmentId"`
}

func GetStudents(students *[]Student) (err error) {
	if err = DB.Find(students).Error; err != nil {
		fmt.Println("Error in GetStudent")
		return err
	}
	return nil
}

func SaveStudents(students []Student) (err error) {
	for _, student := range students {
		if err = DB.Create(&student).Error; err != nil {
			fmt.Println("Error in SaveStudent")
			return err
		}
	}
	return nil
}

func GetStudentsWithDepartment(students *[]StudentsWithDepartment) (err error) {
	if err = DB.Raw("SELECT S.*, D.department_name FROM student AS S INNER JOIN class AS C ON S.class_id = C.id INNER JOIN department AS D ON C.department_id = D.id").Scan(students).Error; err != nil {
		fmt.Println("Error in GetStudentsWithDepartment")
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

func SaveTeachers(teachers []Teacher) (err error) {
	for _, teacher := range teachers {
		if err = DB.Create(&teacher).Error; err != nil {
			fmt.Println("error In SaveTeachers")
			return err
		}
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

func GetTeachersByDepartment(teachers *[]Teacher, departmentID int) (err error) {
	if err = DB.Joins("Class").Find(teachers, fmt.Sprintf("department_id=%d", departmentID)).Error; err != nil {
		fmt.Println("Error in GetTeacherByDepartment")
		return err
	}
	return nil
}

func GetStudentsByDepartment(students *[]Student, departmentID int) (err error) {
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
