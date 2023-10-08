package excel

/*
import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ParseData(env, ns string, bodyBytes []byte) (string, error) {
	var resp Response
	if err := json.Unmarshal(bodyBytes, &resp); err != nil {
		return "", fmt.Errorf("解析 JSON 数据出错: %v", err)
	}

	if len(resp.Data.Result) == 0 {
		return "", fmt.Errorf("cce集群: %s,namespace: %s,无 pods。", env, ns)
	}

	valueSlice := resp.Data.Result[0].Value
	if len(valueSlice) < 2 {
		return "", fmt.Errorf("值列表长度不足")
	}

	valueStr, ok := valueSlice[1].(string)
	if !ok {
		return "", fmt.Errorf("无法转换为字符串")
	}

	valueFloat, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return "", fmt.Errorf("无法转换为 float64: %v", err)
	}

	valueRounded := math.Round(valueFloat*100) / 100
	valueString := fmt.Sprintf("%.2f", valueRounded)

	return valueString, nil
}

func WriteToExcel(filePath, sheetName string, value map[string]string, startRow int) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("打开文件失败：%v", err)
	}

	for col, data := range value {
		cell := fmt.Sprintf("%s%d", col, startRow)
		f.SetCellValue(sheetName, cell, data)
	}

	if err = f.Save(); err != nil {
		return fmt.Errorf("保存文件失败：%v", err)
	}

	return nil
}

func GetSheetAndRangeByBusiness(env string) (string, map[string]string, error) {
	switch env {
	case "hwc-sh1-dev-cluster":
		return "开发业务信息",
			//map[string]string{"A2": "namespace", "B2": "CPU请求(core)", "C2": "内存请求(GB)", "D2": "CPU上限(core)", "E2": "内存上限(GB)", "F2": "无状态服务数量", "G2": "有状态服务数量", "H2": "POD数量", "I2": "POD重启统计（重启次数）"},
			map[string]string{"A2": "namespace", "B2": "CPU request"},
			nil
	case "hwc-sh1-test-cluster":
		return "测试业务信息",
			map[string]string{"A2": "namespace", "B2": "CPU request"},
			nil
	case "hwc-sh1-pre-cluster":
		return "开发业务信息",
			map[string]string{"A2": "namespace", "B2": "CPU request"},
			nil
	case "hwc-sh1-prod-cluster":
		return "测试业务信息",
			map[string]string{"A2": "namespace", "B2": "CPU request"},
			nil
	default:
		return "", map[string]string{}, fmt.Errorf("找不到对应cce集群")

	}
}

func WriteBusinessInfoToExcel(filePath, env, ns string, values string) error {
	startRow := 2 // 设置开始行

	sheetName, mapping, err := GetSheetAndRangeByBusiness(env)

	//fmt.Printf("sheetName: %s,m mapping: %s", sheetName, mapping)
	if err != nil {
		return fmt.Errorf("根据 env 获取 sheetName 错误：%v", err)
	}

	var response Response
	err = json.Unmarshal([]byte(values), &response)
	if err != nil {
		return fmt.Errorf("解析JSON数据失败：%v", err)
	}

	fmt.Printf("\nresponse.Data.Result: %s", response.Data.Result)
	//WriteToExcel(filePath, sheetName, values, startRow)
	for _, result := range response.Data.Result { // 遍历所有的结果

		valueStr, ok := result.Value[1].(string) // 假设第二个值是 CPU 请求
		fmt.Printf("\nvalueStr:%s", valueStr)
		if !ok {
			continue // 如果无法转换为字符串，则跳过这个结果
		}

		dataToWrite := make(map[string]string, len(mapping))

		for cell, field := range mapping {
			switch field {
			case "namespace":
				dataToWrite[cell] = ns
			case "CPU request":
				dataToWrite[cell] = valueStr
			}

		}
		// 写入这个namespace下这个pod的数据，并将startRow增加1以便下次写入下一行
		if err = WriteToExcel(filePath, sheetName, dataToWrite, startRow); err != nil {
			return fmt.Errorf("写入Excel错误：%v", err)
		}

		startRow++
	}
	return nil

}
*/
