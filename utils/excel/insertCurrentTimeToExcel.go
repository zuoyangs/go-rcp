package excel

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

func InsertCurrentTimeToExcel(dst, sheetName string) {
	f, err := excelize.OpenFile(dst)
	if err != nil {
		fmt.Printf("failed to open file: %v", err)
		return
	}
	defer f.Close()

	err = f.InsertRows(sheetName, 1, 1)
	if err != nil {
		fmt.Printf("failed to insert row: %v", err)
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	err = f.SetCellValue(sheetName, "A1", currentTime)
	if err != nil {
		fmt.Printf("failed to set cell value: %v", err)
		return
	}

	err = f.SaveAs(dst)
	if err != nil {
		fmt.Printf("failed to save file: %v", err)
		return
	}
}
