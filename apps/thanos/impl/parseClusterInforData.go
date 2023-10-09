package impl

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ClusterInfo struct {
	ClusterName     string  // 集群名称
	Environment     string  // 环境
	NodeNum         int     // 节点数
	TotalCPU        int     // CPU总额（core）
	TotalMemory     float64 // 内存总额（GB）
	AllocatedCPU    float64 // 已分配的CPU（核心）
	AllocatedMemory float64 // 已分配的内存（GB）
}

// sheetName: "集群信息"
func ParseClusterInfoData(dst, env, ns, key string, jsonStr []byte) interface{} {
	clusterInfo_valueMap := make(map[string]ClusterInfo)
	var resp Response
	if err := json.Unmarshal(jsonStr, &resp); err != nil {
		fmt.Println("错误:", err)
		return nil
	}

	if key == "已分配CPU(core)" {
		startrow := 2
		cpuStr, ok := resp.Data.Result[0].Value[1].(string)
		if !ok {
			fmt.Println("Error: cannot convert Value to string")
			return nil
		}
		cpu, err := strconv.ParseFloat(cpuStr, 64)
		if err != nil {
			fmt.Println("Error: cannot convert CPU value to float")
			return nil
		}

		clusterInfo_valueMap[env] = ClusterInfo{
			ClusterName:  env,
			AllocatedCPU: cpu,
		}
		WriteToExcel_ClusterInfo(dst, "集群信息", env, ns, key, startrow, clusterInfo_valueMap)

	} else if key == "已分配内存(GB)" {
		startrow := 2
		memoryStr, ok := resp.Data.Result[0].Value[1].(string)
		if !ok {
			fmt.Println("Error: cannot convert Value to string")
			return nil
		}
		memory, err := strconv.ParseFloat(memoryStr, 64)
		if err != nil {
			fmt.Println("Error: cannot convert Memory value to float")
			return nil
		}

		clusterInfo_valueMap[env] = ClusterInfo{
			ClusterName:     env,
			AllocatedMemory: memory,
		}
		WriteToExcel_ClusterInfo(dst, "集群信息", env, ns, key, startrow, clusterInfo_valueMap)
	}
	return nil
}

func WriteToExcel_ClusterInfo(dst, sheetName, env, ns, key string, startrow int, clusterInfo_valueMap map[string]ClusterInfo) error {
	f, err := excelize.OpenFile(dst)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	info, ok := clusterInfo_valueMap[env]
	if !ok {
		return fmt.Errorf("no Info for environment: %s", env)
	}

	var cell string
	switch env {
	case "hwc-sh1-dev-cluster":
		cell = "7"
	case "hwc-sh1-test-cluster":
		cell = "6"
	case "hwc-sh1-pre-cluster":
		cell = "5"
	case "hwc-sh1-prod-cluster":
		cell = "4"
	case "hwc-sh1-pub-cluster":
		cell = "3"
	default:
		return fmt.Errorf("unexpected environment: %s", env)
	}

	switch key {
	case "已分配CPU(core)":
		cell = fmt.Sprintf("F%s", cell)
		valueString := strconv.FormatFloat(info.AllocatedCPU, 'f', 2, 64)
		f.SetCellValue(sheetName, cell, valueString)
	case "已分配内存(GB)":
		cell = fmt.Sprintf("G%s", cell)
		valueString := strconv.FormatFloat(info.AllocatedMemory, 'f', 2, 64)
		f.SetCellValue(sheetName, cell, valueString)

	default:
		return fmt.Errorf("unexpected key: %s", key)
	}

	if err = f.Save(); err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	f.Close()
	return nil
}
