package impl

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type BusinessInfo struct {
	ClusterName      string  // 集群名称
	NameSpace        string  // 命名空间
	CPURequest       float64 // CPU请求（核心）
	MemoryRequest    float64 // 内存请求（GB）
	CPULimit         float64 // CPU上限（核心）
	MemoryLimit      float64 // 内存上限（GB）
	DeploymentCount  int     // 无状态服务数量
	StatefulsetCount int     // 有状态服务数量
	PodNum           float64 // POD数量
	PodRestartCount  string  //重启统计(重启次数)
	podRestartName   string  //重启的pod名称

}

// sheetName: "业务信息"
func ParseBusinessInfoData(dst, env, ns, key string, jsonStr []byte, startrow *int) interface{} {
	businessInfo_valueMap := make(map[int]BusinessInfo)
	var resp Response
	if err := json.Unmarshal(jsonStr, &resp); err != nil {
		fmt.Println("错误:", err)
		return nil
	}

	if key == "POD重启统计" {
		var num int = 0
		for _, result := range resp.Data.Result {
			podRestartNameStr := result.Metric.Pod
			podRestartCountStr := result.Value[1].(string)
			podRestartCountInt, err := strconv.Atoi(podRestartCountStr)

			if err != nil {
				fmt.Printf("Failed to convert string %s to integer: %v\n", podRestartCountStr, err)
				continue
			}

			if podRestartCountInt > 0 {
				podRestartStr := fmt.Sprintf("pod名称:%s,重启次数%d", podRestartNameStr, podRestartCountInt)
				businessInfo_valueMap[num] = BusinessInfo{
					ClusterName:     env,
					NameSpace:       ns,
					PodRestartCount: podRestartStr,
					podRestartName:  podRestartNameStr,
				}
				num += 1
				*startrow += 1
				//fmt.Printf("CCE集群名称:%s, namespae:%s,\n 计算指标:%s, Pod名称:%s, 重启次数:%d, excel 表J%d行\n", env, ns, key, podRestartNameStr, podRestartCountInt, *startrow)
			}
		}
		WriteToExcel_BusinessInfo(dst, key, startrow, businessInfo_valueMap[env])
		for i := 0; i < len(businessInfo_valueMap); i++ {
			if businessInfo_valueMap[i].PodRestartCount != "" {
				fmt.Printf("\nCCE集群名称: %s, NameSpace:%s, podRestartName:%s,PodRestartCount:%s\n", businessInfo_valueMap[i].ClusterName, businessInfo_valueMap[i].NameSpace, businessInfo_valueMap[i].podRestartName, businessInfo_valueMap[i].PodRestartCount)

			}
		}
	}
	return nil

}


 // func WriteToExcel_BusinessInfo(dst, env, ns, key string, startrow *int, businessInfomation_valueMap map[string]BusinessInfo) error {
func WriteToExcel_BusinessInfo(dst, key string, startrow *int, businessInfo_valueMap BusinessInfo) error {

	f, err := excelize.OpenFile(dst)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

 	var cell string
	   	var sheetName string
	   	switch businessInfo_valueMap.ClusterName {
	   	case "hwc-sh1-dev-cluster":
	   		sheetName = "开发业务信息"
	   	case "hwc-sh1-test-cluster":
	   		sheetName = "测试业务信息"
	   	case "hwc-sh1-pre-cluster":
	   		sheetName = "预生产业务信息"
	   	case "hwc-sh1-prod-cluster":
	   		sheetName = "生产业务信息"
	   	default:
	   		return fmt.Errorf("unexpected environment: %s", businessInfo_valueMap.ClusterName)
	   	}

	switch key {
		case "CPU请求(core)":
	   		cell = fmt.Sprintf("F%s", cell)
	   		valueString := strconv.FormatFloat(info.CPURequest, 'f', 2, 64)
	   		f.SetCellValue(sheetName, cell, valueString)
	   	case "内存请求(GB)":
	   		cell = fmt.Sprintf("G%s", cell)
	   		valueString := strconv.FormatFloat(info.MemoryRequest, 'f', 2, 64)
	   		f.SetCellValue(sheetName, cell, valueString)
	   	case "CPU上限(core)":
	   		cell = fmt.Sprintf("F%s", cell)
	   		valueString := strconv.FormatFloat(info.CPULimit, 'd', 2, 64)
	   		f.SetCellValue(sheetName, cell, valueString)
	   	case "内存上限(GB)":
	   		cell = fmt.Sprintf("G%s", cell)
	   		valueString := strconv.FormatFloat(info.MemoryLimit, 'e', 2, 64)
	   		f.SetCellValue(sheetName, cell, valueString)
	   	case "无状态服务数量":
	   		cell = fmt.Sprintf("F%s", cell)
	   		valueString := strconv.Itoa(info.DeploymentCount)
	   		f.SetCellValue(sheetName, cell, valueString)
	   	case "有状态服务数量":
	   		cell = fmt.Sprintf("G%s", cell)
	   		valueString := strconv.Itoa(info.StatefulsetCount)
	   		f.SetCellValue(sheetName, cell, valueString)
	   	case "POD数量":
	   		cell = fmt.Sprintf("F%s", cell)
	   		valueString := strconv.FormatFloat(info.PodNum, 'h', 2, 64)
	   		f.SetCellValue(sheetName, cell, valueString)
	case "POD重启统计":
		fmt.Printf("\n businessInfo_valueMap:%s\n", businessInfo_valueMap)
		// 使用 sort.Slice 函数对切片进行降序排序
		sort.Slice(businessInfo_valueMap, func(i, j int) bool {
			return businessInfo_valueMap[i].PodRestartCount > businessInfo_valueMap[j].PodRestartCount
		})

		// 打印排序后的结果
		for _, info := range businessInfo_valueMap {
			fmt.Printf("ClusterName: %s, NameSpace: %s, PodRestartCount: %d, podRestartName: %s\n", info.ClusterName, info.NameSpace, info.PodRestartCount, info.podRestartName)
		}
		 		for _, result := range businessInfo_valueMap.podRestartName {
			if BusinessInfo_valueMap.PodRestartCount
			cell = fmt.Sprintf("C%d", *startrow)
			f.SetCellValue(sheetName, cell, info.PodRestartCount)
			Cluster_cell := fmt.Sprintf("A%d", *startrow)
			f.SetCellValue(sheetName, Cluster_cell, info.ClusterName)
			NameSpace_cell := fmt.Sprintf("B%d", *startrow)
			f.SetCellValue(sheetName, NameSpace_cell, info.NameSpace)
		}
	default:
		return fmt.Errorf("unexpected key: %s", key)

	}

	if err = f.Save(); err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	f.Close()
	return nil
}
