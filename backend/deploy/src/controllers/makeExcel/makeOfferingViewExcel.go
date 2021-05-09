package makeExcel

import (
	"fmt"
	"strings"
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func SaveOfferingViewExcel(date time.Time) {
	f := excelize.NewFile()
	var offeringDiarys []models.OfferingDiary
	var createdBys []string
	models.GetOfferingDiaryByDate(&offeringDiarys, date)

	for i := 0; i < len(offeringDiarys); i++ {
		hasSameName := false
		for j := i + 1; j < len(offeringDiarys); j++ {
			if offeringDiarys[i].CreatedBy == offeringDiarys[j].CreatedBy {
				hasSameName = true
			}
		}

		if !hasSameName {
			createdBys = append(createdBys, offeringDiarys[i].CreatedBy)
		}
	}
	offeringDiarySheetName := date.Format("2006-01-12") + "_헌금통계표"
	index := f.NewSheet(offeringDiarySheetName)

	f.SetCellValue(offeringDiarySheetName, "A1", offeringDiarySheetName)
	f.SetCellValue(offeringDiarySheetName, "A2", fmt.Sprintf("작성자: %s", strings.Join(createdBys, ",")))
	f.SetCellValue(offeringDiarySheetName, "A3", ("열람 날짜 " + date.Format("2006-01-02")))
	f.MergeCell(offeringDiarySheetName, "A1", "D1")
	f.MergeCell(offeringDiarySheetName, "A2", "D2")
	f.MergeCell(offeringDiarySheetName, "A3", "D3")

	f.SetCellValue(offeringDiarySheetName, "B4", "1부")
	f.SetCellValue(offeringDiarySheetName, "C4", "2부")
	f.SetCellValue(offeringDiarySheetName, "D4", "합계")

	f.SetCellValue(offeringDiarySheetName, "A5", "주일헌금")
	f.SetCellValue(offeringDiarySheetName, "A6", "십일조")
	f.SetCellValue(offeringDiarySheetName, "A7", "감사헌금")
	f.SetCellValue(offeringDiarySheetName, "A8", "절기헌금")
	f.SetCellValue(offeringDiarySheetName, "A9", "기타헌금")
	f.SetCellValue(offeringDiarySheetName, "A10", "합계")

	// 원하는 날짜

	var (
		totalWeekOfferingCost     []int = []int{0, 0}
		totalTithOfferingCost     []int = []int{0, 0}
		totalThanksOfferingCost   []int = []int{0, 0}
		totalSeasonalOfferingCost []int = []int{0, 0}
		totalEtcOfferingCost      []int = []int{0, 0}

		// [departmentID - 1 ][name]
		tithOfferingName     map[int]string = make(map[int]string)
		thanksOfferingName   map[int]string = make(map[int]string)
		seasonalOfferingName map[int]string = make(map[int]string)
		etcOfferingName      map[int]string = make(map[int]string)
	)

	for _, offeringDiary := range offeringDiarys {
		switch offeringDiary.OfferingType.OfferingTypeName {
		case "주일헌금":
			if offeringDiary.DepartmentID == 1 {
				totalWeekOfferingCost[0] += offeringDiary.Cost
			}
			if offeringDiary.DepartmentID == 2 {
				totalWeekOfferingCost[1] += offeringDiary.Cost
			}

		case "십일조헌금":
			if offeringDiary.DepartmentID == 1 {
				totalTithOfferingCost[0] += offeringDiary.Cost
				tithOfferingName[0] += offeringDiary.Student.Name + " "
			}
			if offeringDiary.DepartmentID == 2 {
				totalTithOfferingCost[1] += offeringDiary.Cost
				tithOfferingName[1] += offeringDiary.Student.Name + " "
			}

		case "감사헌금":
			if offeringDiary.DepartmentID == 1 {
				totalThanksOfferingCost[0] += offeringDiary.Cost
				thanksOfferingName[0] += offeringDiary.Student.Name + " "
			}
			if offeringDiary.DepartmentID == 2 {
				totalThanksOfferingCost[1] += offeringDiary.Cost
				thanksOfferingName[1] += offeringDiary.Student.Name + " "
			}

		case "절기헌금":
			if offeringDiary.DepartmentID == 1 {
				totalSeasonalOfferingCost[0] += offeringDiary.Cost
				seasonalOfferingName[0] += offeringDiary.Student.Name + " "
			}
			if offeringDiary.DepartmentID == 2 {
				totalSeasonalOfferingCost[1] += offeringDiary.Cost
				seasonalOfferingName[1] += offeringDiary.Student.Name + " "
			}

		case "기타헌금":
			if offeringDiary.DepartmentID == 1 {
				totalEtcOfferingCost[0] += offeringDiary.Cost
				etcOfferingName[0] += offeringDiary.Student.Name + " "
			}
			if offeringDiary.DepartmentID == 2 {
				totalEtcOfferingCost[1] += offeringDiary.Cost
				etcOfferingName[1] += offeringDiary.Student.Name + " "
			}
		}
	}

	f.SetCellValue(offeringDiarySheetName, "B5", totalWeekOfferingCost[0])
	f.SetCellValue(offeringDiarySheetName, "C5", totalWeekOfferingCost[1])
	f.SetCellValue(offeringDiarySheetName, "D5", totalWeekOfferingCost[0]+totalWeekOfferingCost[1])

	f.SetCellValue(offeringDiarySheetName, "B6", totalTithOfferingCost[0])
	f.SetCellValue(offeringDiarySheetName, "C6", totalTithOfferingCost[1])
	f.SetCellValue(offeringDiarySheetName, "D6", totalTithOfferingCost[0]+totalTithOfferingCost[1])

	f.SetCellValue(offeringDiarySheetName, "B7", totalThanksOfferingCost[0])
	f.SetCellValue(offeringDiarySheetName, "C7", totalThanksOfferingCost[1])
	f.SetCellValue(offeringDiarySheetName, "D7", totalThanksOfferingCost[0]+totalThanksOfferingCost[1])

	f.SetCellValue(offeringDiarySheetName, "B8", totalSeasonalOfferingCost[0])
	f.SetCellValue(offeringDiarySheetName, "C8", totalSeasonalOfferingCost[1])
	f.SetCellValue(offeringDiarySheetName, "D8", totalSeasonalOfferingCost[0]+totalSeasonalOfferingCost[1])

	f.SetCellValue(offeringDiarySheetName, "B9", totalEtcOfferingCost[0])
	f.SetCellValue(offeringDiarySheetName, "C9", totalEtcOfferingCost[1])
	f.SetCellValue(offeringDiarySheetName, "D9", totalEtcOfferingCost[0]+totalEtcOfferingCost[1])

	var totalOfferingCostByDepartment []int = []int{0, 0}
	totalOfferingCostByDepartment[0] = totalWeekOfferingCost[0] + totalTithOfferingCost[0] + totalThanksOfferingCost[0] + totalSeasonalOfferingCost[0] + totalEtcOfferingCost[0]
	totalOfferingCostByDepartment[1] = totalWeekOfferingCost[1] + totalTithOfferingCost[1] + totalThanksOfferingCost[1] + totalSeasonalOfferingCost[1] + totalEtcOfferingCost[1]
	f.SetCellValue(offeringDiarySheetName, "B10", totalOfferingCostByDepartment[0])
	f.SetCellValue(offeringDiarySheetName, "C10", totalOfferingCostByDepartment[1])
	f.SetCellValue(offeringDiarySheetName, "D10", totalOfferingCostByDepartment[0]+totalOfferingCostByDepartment[1])

	f.MergeCell(offeringDiarySheetName, "A12", "D12")
	f.MergeCell(offeringDiarySheetName, "A13", "D13")
	f.MergeCell(offeringDiarySheetName, "A14", "D14")
	f.MergeCell(offeringDiarySheetName, "A15", "D15")

	f.SetCellValue(offeringDiarySheetName, "A12", fmt.Sprintf("1부: %s\n2부: %s", tithOfferingName[0], tithOfferingName[1]))
	f.SetCellValue(offeringDiarySheetName, "A13", fmt.Sprintf("1부: %s\n2부: %s", thanksOfferingName[0], thanksOfferingName[1]))
	f.SetCellValue(offeringDiarySheetName, "A14", fmt.Sprintf("1부: %s\n2부: %s", seasonalOfferingName[0], seasonalOfferingName[1]))
	f.SetCellValue(offeringDiarySheetName, "A15", fmt.Sprintf("1부: %s\n2부: %s", etcOfferingName[0], etcOfferingName[1]))

	// 주일헌금 + 십일조 + 감사헌금 + 절기헌금 + 기타헌금 / 합계

	f.SetActiveSheet(index)

	if err := f.SaveAs("data/offeringDiary/excel/" + date.Format("2006-01-02") + ".xlsx"); err != nil {
		fmt.Println("error in saving excel: ", err)
	}
}
