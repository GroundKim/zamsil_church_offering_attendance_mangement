package models

type Department struct {
	ID           int
	DepartmentID int `gorm:"not null;"`

	Classes []Class `gorm:"foreignKeys:DepartmentID"`
}

type Class struct {
	ID           int
	ClassID      int `gorm:"not null;"`
	DepartmentID int `gorm:"not null;"`
}

func (Department) TableName() string {
	return "department"
}

func (Class) TableName() string {
	return "class"
}
