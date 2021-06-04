package models

// member

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Student struct {
	ID                int        `json:"studentId"`
	Name              string     `gorm:"not null;" json:"name"`
	ClassID           int        `gorm:"not null;" json:"classId"`
	CreatedAt         *time.Time `gorm:"not null; default:current_timestamp(3);"  json:"createdAt"`
	UpdatedAt         *time.Time `gorm:"null; default:null" json:"updatedAt"`
	DayOfBirth        *time.Time `gorm:"type:date; null;" json:"dayOfBirth"`
	Address           *string    `json:"address"`
	PhoneNumber       *string    `json:"phoneNumber"`
	ParentPhoneNumber *string    `json:"parentPhoneNumber"`
	SchoolName        *string    `json:"schoolName"`

	Class Class `gorm:"references:ID" json:"-"`
}

type Teacher struct {
	ID          int       `json:"teacherId"`
	Name        string    `gorm:"not null;" json:"name"`
	ClassID     int       `gorm:"not null;" json:"classId"`
	CreatedAt   time.Time `gorm:"not null;" json:"createdAt"`
	PhoneNumber *string   `gorm:"null" json:"phoneNumber"`

	Class Class `gorm:"references:ID" json:"-"`
}

type StudentsWithDepartment struct {
	ID             int       `json:"studentId"`
	Name           string    `json:"name"`
	ClassID        int       `json:"classId"`
	CreatedAt      time.Time `json:"createdAt"`
	DepartmentName int       `json:"departmentId"`
}

func GetStudents(students *[]Student) (err error) {
	if err = DB.Preload("Class").Find(&students).Error; err != nil {
		fmt.Println("Error in GetStudent")
		return err
	}
	return nil
}

func SaveStudents(students *[]Student) (err error) {
	if err = DB.Create(&students).Error; err != nil {
		fmt.Println("Error in SaveStudent")
		return err
	}
	// log
	for _, student := range *students {
		savedStudent, _ := json.Marshal(student)
		AuthToken.User.ActiveStamp("save student", string(savedStudent))
	}
	return nil
}

// need maintenance when add field in student
func PutStudent(student Student) (err error) {
	var oldStudent Student
	DB.First(&oldStudent, student.ID)

	// avoid to get nil pointer
	fieldNames := []string{"DayOfBirth", "Address", "PhoneNumber", "ParentPhoneNumber", "SchoolName"}
	values := []string{}
	v := reflect.ValueOf(student)
	for _, fieldName := range fieldNames {
		if v.FieldByName(fieldName).IsNil() {
			values = append(values, "null")
		} else {
			add := v.FieldByName(fieldName).Interface()
			value := reflect.Indirect(reflect.ValueOf(add)).String()
			values = append(values, "\""+value+"\"")
		}
	}

	updateQuery := fmt.Sprintf("UPDATE student SET updated_at = '%s', name = '%s', class_id = '%d', day_of_birth = %s, address = %s, phone_number = %s, parent_phone_number = %s, school_name = %s WHERE id = %d",
		// it does not look good
		time.Now().Format("2006-01-02 15:04:05"), student.Name, student.ClassID, values[0], values[1], values[2], values[3], values[4], student.ID)

	if err = DB.Exec(updateQuery).Error; err != nil {
		return err
	}
	// user log
	putStudent, _ := json.Marshal(student)
	oldStudentJSON, _ := json.Marshal(oldStudent)
	AuthToken.User.ActiveStamp("PUT", string(oldStudentJSON)+"->"+string(putStudent))
	return nil
}

func DeleteStudent(student *Student) (err error) {
	if err = DB.Find(&student).Error; err != nil {
		return err
	}

	if err = DB.Delete(&student).Error; err != nil {
		return err
	}

	// user log
	deleteStudent, _ := json.Marshal(student)
	AuthToken.User.ActiveStamp("DELETE", string(deleteStudent))
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
	if err = DB.Joins("Class").Where("department_id = ? AND Class.name = ?", departmentID, className).Find(teachers).Error; err != nil {
		fmt.Println("Error in GetTeacherByClassName")
		return err
	}
	return nil
}

func GetStudentByClassNameAndDepartment(students *[]Student, departmentID int, className string) (err error) {
	if err = DB.Joins("Class").Where("department_id = ? AND Class.name = ?", departmentID, className).Find(students).Error; err != nil {
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
