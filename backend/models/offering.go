package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type OfferingDiary struct {
	ID             int       `json:"offeringDiaryId"`
	StudentID      *int      `json:"studentId"`
	OfferingTypeID int       `gorm:"not null;" json:"offeringTypeId"`
	DepartmentID   int       `gorm:"not null;" json:"departmentId"`
	Cost           int       `gorm:"not null;" json:"offeringCost"`
	OfferedAt      time.Time `gorm:"not null; type:date" json:"offeredAt"`
	CreatedAt      time.Time `gorm:"not null;" json:"createdAt"`
	CreatedBy      string    `gorm:"not null;" json:"createdBy"`

	Student      Student      `gorm:"references:ID" json:"student"`
	OfferingType OfferingType `gorm:"references:ID" json:"offeringType"`
	Department   Department   `gorm:"references:ID" json:"department"`
}

type OfferingType struct {
	ID               int    `json:"offeringTypeId"`
	OfferingTypeName string `gorm:"not null;" json:"offeringTypeName"`
}

func GetOfferingDiaryByDate(offeringDiaries *[]OfferingDiary, date time.Time) (err error) {
	theDay := date.Format("2006-01-02 ") + "00:00:00"
	theDayRange := theDay[0:11] + "23:59:59"
	if err = DB.Preload("OfferingType").Preload("Student").Preload("Department").Preload("Student.Class").Where("offered_at BETWEEN ? AND ?", theDay, theDayRange).Find(&offeringDiaries).Error; err != nil {
		fmt.Println("Error in get Offering Diary by date")
		return err
	}
	return nil
}

func GetOfferingDiaryByYear(offeringDiaries *[]OfferingDiary, date time.Time) (err error) {
	theYear := date.Format("2006-01-02 ") + "00:00:00"
	theYearRange := theYear[0:4] + "-12-31 11:59:59"
	if err = DB.Preload("OfferingType").Preload("Student").Where("offered_at BETWEEN ? AND ?", theYear, theYearRange).Find(&offeringDiaries).Error; err != nil {
		fmt.Println("Error in get Offering Diary by date")
		return err
	}
	return nil

}

func (OfferingDiary *OfferingDiary) SaveOfferingDiary() (err error) {
	if err = DB.Create(&OfferingDiary).Error; err != nil {
		fmt.Println("Error in create Offering Diary")
		return err
	}
	return nil
}

func GetOfferingType(OfferingTypes *[]OfferingType) (err error) {
	if err = DB.Find(OfferingTypes).Error; err != nil {
		fmt.Println("Error in GetOfferingType")
		return err
	}
	return nil
}

func ChangeOfferingDiaryTypeAndCost(offeringDiary *OfferingDiary) (err error) {
	var oldOfferingDiary OfferingDiary
	DB.Find(&oldOfferingDiary, offeringDiary.ID)
	offeringDiary.CreatedAt = oldOfferingDiary.CreatedAt
	offeringDiary.CreatedBy = oldOfferingDiary.CreatedBy
	offeringDiary.OfferedAt = oldOfferingDiary.OfferedAt
	offeringDiary.StudentID = &oldOfferingDiary.ID
	offeringDiary.DepartmentID = oldOfferingDiary.DepartmentID

	if err = DB.Save(&offeringDiary).Error; err != nil {
		return err
	}

	marshaledNewOfferingDiary, _ := json.Marshal(offeringDiary)
	marshaledOldOfferingDiary, _ := json.Marshal(oldOfferingDiary)
	AuthToken.User.ActiveStamp("PUT OFFERING DIARY", string(marshaledNewOfferingDiary)+"->"+string(marshaledOldOfferingDiary))
	return nil
}

func DeleteOfferingDiary(offeringDiary *OfferingDiary) (err error) {
	if err = DB.Preload("Student").Find(&offeringDiary).Error; err != nil {
		return err
	}

	if err = DB.Delete(&offeringDiary).Error; err != nil {
		return err
	}

	// user log
	deletedOfferingDiary, _ := json.Marshal(offeringDiary)
	AuthToken.User.ActiveStamp("DELETE OFFERING DIARY", string(deletedOfferingDiary))
	return nil
}

// handling offering summary
type OfferingCost struct {
	TotalWeekOfferingCost     []int    `json:"totalWeekOfferingCost"`
	TotalTitheOfferingCost    []int    `json:"totalTitheOfferingCost"`
	TotalThanksOfferingCost   []int    `json:"totalThanksOfferingCost"`
	TotalSeasonalOfferingCost []int    `json:"totalSeasonalOfferingCost"`
	TotalEtcOfferingCost      []int    `json:"totalEtcOfferingCost"`
	CreatedBy                 []string `json:"createdBy"`
}

type OfferingCostWrapper struct {
	OfferingType OfferingType `json:"offeringType"`
	OfferingCost []int        `json:"offeringCost"`
}

// need maintenance if offering type get to be changed
func GetSummaryOfferingCostByDate(date time.Time) (offeringCostWrappers []OfferingCostWrapper, err error) {
	offeringCosts := &OfferingCost{}

	offeringCosts.TotalWeekOfferingCost = []int{0, 0}
	offeringCosts.TotalTitheOfferingCost = []int{0, 0}
	offeringCosts.TotalThanksOfferingCost = []int{0, 0}
	offeringCosts.TotalSeasonalOfferingCost = []int{0, 0}
	offeringCosts.TotalEtcOfferingCost = []int{0, 0}

	var offeringDiaries []OfferingDiary
	var createdBys []string
	err = GetOfferingDiaryByDate(&offeringDiaries, date)

	for i := 0; i < len(offeringDiaries); i++ {
		hasSameName := false
		for j := i + 1; j < len(offeringDiaries); j++ {
			if offeringDiaries[i].CreatedBy == offeringDiaries[j].CreatedBy {
				hasSameName = true
			}
		}

		if !hasSameName {
			createdBys = append(createdBys, offeringDiaries[i].CreatedBy)
		}
	}
	offeringCosts.CreatedBy = createdBys

	for _, offeringDiary := range offeringDiaries {
		switch offeringDiary.OfferingType.OfferingTypeName {
		case "주일헌금":
			if offeringDiary.DepartmentID == 1 {
				offeringCosts.TotalWeekOfferingCost[0] += offeringDiary.Cost
			}
			if offeringDiary.DepartmentID == 2 {
				offeringCosts.TotalWeekOfferingCost[1] += offeringDiary.Cost
			}

		case "십일조헌금":
			if offeringDiary.DepartmentID == 1 {
				offeringCosts.TotalTitheOfferingCost[0] += offeringDiary.Cost
			}
			if offeringDiary.DepartmentID == 2 {
				offeringCosts.TotalTitheOfferingCost[1] += offeringDiary.Cost
			}

		case "감사헌금":
			if offeringDiary.DepartmentID == 1 {
				offeringCosts.TotalThanksOfferingCost[0] += offeringDiary.Cost
			}
			if offeringDiary.DepartmentID == 2 {
				offeringCosts.TotalThanksOfferingCost[1] += offeringDiary.Cost
			}

		case "절기헌금":
			if offeringDiary.DepartmentID == 1 {
				offeringCosts.TotalSeasonalOfferingCost[0] += offeringDiary.Cost
			}
			if offeringDiary.DepartmentID == 2 {
				offeringCosts.TotalSeasonalOfferingCost[1] += offeringDiary.Cost
			}

		case "기타헌금":
			if offeringDiary.DepartmentID == 1 {
				offeringCosts.TotalEtcOfferingCost[0] += offeringDiary.Cost
			}
			if offeringDiary.DepartmentID == 2 {
				offeringCosts.TotalEtcOfferingCost[1] += offeringDiary.Cost
			}
		}
	}

	weekOfferingCostWrapper := &OfferingCostWrapper{OfferingType: OfferingType{ID: 1, OfferingTypeName: "주일헌금"}, OfferingCost: offeringCosts.TotalWeekOfferingCost}
	titheOfferingCostWrapper := &OfferingCostWrapper{OfferingType: OfferingType{ID: 2, OfferingTypeName: "십일조헌금"}, OfferingCost: offeringCosts.TotalTitheOfferingCost}
	thanksOfferingCostWrapper := &OfferingCostWrapper{OfferingType: OfferingType{ID: 3, OfferingTypeName: "감사헌금"}, OfferingCost: offeringCosts.TotalThanksOfferingCost}
	seasonalOfferingCostWrapper := &OfferingCostWrapper{OfferingType: OfferingType{ID: 4, OfferingTypeName: "절기헌금"}, OfferingCost: offeringCosts.TotalSeasonalOfferingCost}
	etcOfferingCostWrapper := &OfferingCostWrapper{OfferingType: OfferingType{ID: 5, OfferingTypeName: "기타헌금"}, OfferingCost: offeringCosts.TotalEtcOfferingCost}

	offeringCostWrappers = append(offeringCostWrappers, *weekOfferingCostWrapper, *titheOfferingCostWrapper, *thanksOfferingCostWrapper, *seasonalOfferingCostWrapper, *etcOfferingCostWrapper)
	return offeringCostWrappers, err
}
func (OfferingDiary) TableName() string {
	return "offering_diary"
}

func (OfferingType) TableName() string {
	return "offering_type"
}
