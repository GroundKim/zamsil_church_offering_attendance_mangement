package makeExcel

import (
	"fmt"
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func SaveExcel(date time.Time, createdBy string) {
	f := excelize.NewFile()
	offeringDiaryName := "헌금통계표"

	index := f.NewSheet(offeringDiaryName)

	f.SetCellValue(offeringDiaryName, "A1", offeringDiaryName)
	f.SetCellValue(offeringDiaryName, "A2", ("작성자: " + createdBy))
	f.SetCellValue(offeringDiaryName, "A3", ("열람 날짜 " + time.Now().Format("2006-01-02")))
	f.MergeCell(offeringDiaryName, "A1", "D1")
	f.MergeCell(offeringDiaryName, "A2", "D2")
	f.MergeCell(offeringDiaryName, "A3", "D3")

	f.SetCellValue(offeringDiaryName, "B4", "1부")
	f.SetCellValue(offeringDiaryName, "C4", "2부")
	f.SetCellValue(offeringDiaryName, "D4", "합계")

	f.SetCellValue(offeringDiaryName, "A5", "주일헌금")
	f.SetCellValue(offeringDiaryName, "A6", "십일조")
	f.SetCellValue(offeringDiaryName, "A7", "감사헌금")
	f.SetCellValue(offeringDiaryName, "A8", "절기헌금")

	// 원하는 날짜
	now := time.Now()

	var offeringDiarys []models.OfferingDiary

	models.GetSpecificOfferingDiaryByDate(&offeringDiarys, now)

	// 주일헌금 + 십일조 + 감사헌금 + 절기헌금 + 기타헌금 / 합계

	f.SetActiveSheet(index)

	if err := f.SaveAs("data/offeringDiary/excel/" + date.Format("2006-01-02") + ".xlsx"); err != nil {
		fmt.Println("error in saving excel: ", err)
	}
}
