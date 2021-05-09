package makeExcel

import (
	"fmt"
	"strings"
	"time"
	"zamsil_church_offering_attendance_mangement/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type attendanceInfo struct {
	departmentID int
	className    string
	teacherNames []string
	enrollment   int
	attendance   int
	absence      int
}

func (a *attendanceInfo) CountAttendanceByClass(departmentID int, className string, diaries []models.AttendanceDiary) {
	a.departmentID = departmentID
	a.className = className

	var teachers []models.Teacher
	models.GetTeacherByClassNameAndDepartment(&teachers, departmentID, className)

	var teacherNames []string
	for _, teacher := range teachers {
		teacherNames = append(teacherNames, teacher.Name)
	}
	a.teacherNames = teacherNames

	var students []models.Student
	models.GetStudentByClassNameAndDepartment(&students, departmentID, className)
	a.enrollment = len(students)

	for _, student := range students {
		for _, diary := range diaries {
			if student.Name == diary.Student.Name {
				a.attendance++
			}
		}
	}

	a.absence = a.enrollment - a.attendance
}

func SaveAttendaceViewExcel(date time.Time) {

	// get attendanceDiaries from DB
	var attendanceDiaries []models.AttendanceDiary
	models.GetAttendanceViewByDate(&attendanceDiaries, date)

	var classesDepartmentOne []models.Class
	models.GetClassesByDepartment(&classesDepartmentOne, 1)
	var attendanceInfoDepartmentOne []attendanceInfo

	var classesDepartmentTwo []models.Class
	models.GetClassesByDepartment(&classesDepartmentTwo, 2)
	var attendanceInfoDepartmentTwo []attendanceInfo

	for _, class := range classesDepartmentOne {
		var attendanceInfo attendanceInfo
		attendanceInfo.CountAttendanceByClass(1, class.Name, attendanceDiaries)
		attendanceInfoDepartmentOne = append(attendanceInfoDepartmentOne, attendanceInfo)
	}

	for _, class := range classesDepartmentTwo {
		var attendanceInfo attendanceInfo
		attendanceInfo.CountAttendanceByClass(2, class.Name, attendanceDiaries)
		attendanceInfoDepartmentTwo = append(attendanceInfoDepartmentTwo, attendanceInfo)
	}

	// get createdBys from attendanceDiaries
	var createdBys []string
	for i := 0; i < len(attendanceDiaries); i++ {
		hasSameName := false
		for j := i + 1; j < len(attendanceDiaries); j++ {
			if attendanceDiaries[i].CreatedBy == attendanceDiaries[j].CreatedBy {
				hasSameName = true
				continue
			}
		}

		if !hasSameName {
			createdBys = append(createdBys, attendanceDiaries[i].CreatedBy)
		}
	}

	/* #region EXCEL ui */
	f := excelize.NewFile()

	attendanceDiarySheetName := date.Format("2006-01-02") + "_출석부"
	index := f.NewSheet(attendanceDiarySheetName)

	f.SetActiveSheet(index)

	f.SetCellValue(attendanceDiarySheetName, "A1", attendanceDiarySheetName)

	f.SetCellValue(attendanceDiarySheetName, "A4", fmt.Sprintf("%d년 %d월 %d일", date.Year(), int(date.Month()), date.Day()))

	f.SetCellValue(attendanceDiarySheetName, "A5", "출석통계")
	f.SetCellValue(attendanceDiarySheetName, "A6", "작성자: "+strings.Join(createdBys, ", "))

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

	var lastDepartmetOneColumn int = 0
	var lastDepartmetTwoColumn int = 0
	var totalDepartmentOneAttendance int = 0
	var totalDepartmentTwoAttendance int = 0
	for i, attendance := range attendanceInfoDepartmentOne {
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("A%d", i+8), attendance.className)
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("B%d", i+8), strings.Join(attendance.teacherNames, ", "))
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("C%d", i+8), attendance.enrollment)
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("D%d", i+8), attendance.attendance)
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("E%d", i+8), attendance.absence)
		lastDepartmetOneColumn = i
		totalDepartmentOneAttendance += attendance.attendance
	}

	for i, attendance := range attendanceInfoDepartmentTwo {
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("H%d", i+8), attendance.className)
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("I%d", i+8), strings.Join(attendance.teacherNames, ", "))
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("J%d", i+8), attendance.enrollment)
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("K%d", i+8), attendance.attendance)
		f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("L%d", i+8), attendance.absence)
		lastDepartmetTwoColumn = i
		totalDepartmentTwoAttendance += attendance.attendance
	}
	var lastColumn int = 0
	if lastDepartmetOneColumn > lastDepartmetTwoColumn {
		lastColumn = lastDepartmetOneColumn
	} else {
		lastColumn = lastDepartmetTwoColumn
	}

	f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("A%d", lastColumn+10), "1부 소계: ")
	f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("B%d", lastColumn+10), fmt.Sprintf("%d 명", totalDepartmentOneAttendance))

	f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("H%d", lastColumn+10), "2부 소계: ")
	f.SetCellValue(attendanceDiarySheetName, fmt.Sprintf("I%d", lastColumn+10), fmt.Sprintf("%d 명", totalDepartmentTwoAttendance))
	/* #endregion */

	if err := f.SaveAs("data/attendanceDiary/excel/" + date.Format("2006-01-02") + ".xlsx"); err != nil {
		fmt.Println("error in saving excel: ", err)
	}
}
