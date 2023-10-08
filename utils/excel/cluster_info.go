package excel

/*
func getCellByEnv(env string, column string) (string, error) {
	switch env {
	case "hwc-sh1-dev-cluster":
		return column + "7", nil
	case "hwc-sh1-test-cluster":
		return column + "6", nil
	case "hwc-sh1-pre-cluster":
		return column + "5", nil
	case "hwc-sh1-pub-cluster":
		return column + "3", nil
	case "hwc-sh1-prod-cluster":
		return column + "4", nil
	default:
		return "", fmt.Errorf("找不到对应cce集群")
	}
}

func WriteToExcel_AllocatedCPU(filePath, sheetName, env string, value []byte) error {

	valueString, err := ProcessData(value)
	if err != nil {
		return fmt.Errorf("error processing data: %v", err)
	}

	cell, err := getCellByEnv(env, "F")
	if err != nil {
		return fmt.Errorf("error getting cell by env: %v", err)
	}

	err = writeToExcel(filePath, sheetName, cell, valueString)
	if err != nil {
		return fmt.Errorf("error writing to Excel: %v", err)
	}

	return nil
}

func WriteToExcel_AllocatedMEM(filePath, sheetName, env string, value []byte) error {

	valueString, err := ProcessData(value)
	if err != nil {
		return fmt.Errorf("error processing data: %v", err)
	}

	cell, err := getCellByEnv(env, "G")
	if err != nil {
		return fmt.Errorf("error getting cell by env: %v", err)
	}

	err = writeToExcel(filePath, sheetName, cell, valueString)
	if err != nil {
		return fmt.Errorf("error writing to Excel: %v", err)
	}

	return nil
}
*/
