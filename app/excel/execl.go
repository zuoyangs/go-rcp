package execl

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

func FillExcelData(filePath, sheet, cell, value string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}

	err = f.SetCellValue(sheet, cell, value)
	if err != nil {
		return err
	}

	err = f.Save()
	if err != nil {
		return err
	}

	return nil
}

func WriteToExcel(avgCPUUsage float64, avgMemUsage float64, peakCPUUsage float64, peakMemUsage float64) error {
	f := excelize.NewFile()
	sheetName := "集群信息"

	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}
	f.SetActiveSheet(index)

	f.SetCellValue(sheetName, "F1", "CPU使用率（均值）")
	f.SetCellValue(sheetName, "G1", "内存使用率(均值)")
	f.SetCellValue(sheetName, "H1", "CPU使用率(峰值)")
	f.SetCellValue(sheetName, "I1", "内存使用率（峰值）")

	// 假设从第二行开始写入数据
	rowNumber := 2

	f.SetCellValue(sheetName, "F"+strconv.Itoa(rowNumber), avgCPUUsage)
	f.SetCellValue(sheetName, "G"+strconv.Itoa(rowNumber), avgMemUsage)
	f.SetCellValue(sheetName, "H"+strconv.Itoa(rowNumber), peakCPUUsage)
	f.SetCellValue(sheetName, "I"+strconv.Itoa(rowNumber), peakMemUsage)

	if err := f.SaveAs("集群信息.xlsx"); err != nil {
		return err
	}

	return nil
}
