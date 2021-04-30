package makeExcel

import (
	"fmt"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func SaveAttendaceViewExcel(date time.Time) {
	f := excelize.NewFile()
	attendanceDiarySheetName := date.Format("2006-01-12") + "_출석부"
	index := f.NewSheet(attendanceDiarySheetName)

	f.SetActiveSheet(index)

	f.SetCellValue(attendanceDiarySheetName, "A1", attendanceDiarySheetName)

	f.SetCellValue(attendanceDiarySheetName, "A4", fmt.Sprintf("%d년 %s월 %d일", date.Year(), date.Month().String(), date.Day()))

	f.SetCellValue(attendanceDiarySheetName, "A5", "출석통계")

	f.SetCellValue(attendanceDiarySheetName, "A7", "1부")
	f.SetCellValue(attendanceDiarySheetName, "B7", "교사")
	f.SetCellValue(attendanceDiarySheetName, "C7", "재적")
	f.SetCellValue(attendanceDiarySheetName, "D7", "출석")
	f.SetCellValue(attendanceDiarySheetName, "E7", "결석")
	f.SetCellValue(attendanceDiarySheetName, "F7", "신입")

	f.SetCellValue(attendanceDiarySheetName, "H7", "2부")
	f.SetCellValue(attendanceDiarySheetName, "I7", "교사")
	f.SetCellValue(attendanceDiarySheetName, "J7", "재적")
	f.SetCellValue(attendanceDiarySheetName, "K7", "출석")
	f.SetCellValue(attendanceDiarySheetName, "L7", "결석")
	f.SetCellValue(attendanceDiarySheetName, "M7", "신입")

	// department 1 반 별로, 선생님, 출석수, 결석수

	if err := f.SaveAs("data/attendanceDiary/excel/" + date.Format("2006-01-02") + ".xlsx"); err != nil {
		fmt.Println("error in saving excel: ", err)
	}
}
