package utils

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Response struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name                string `json:"__name__"`
				Cluster             string `json:"cluster"`
				ClusterCategory     string `json:"cluster_category"`
				ClusterName         string `json:"cluster_name"`
				Container           string `json:"container"`
				Endpoint            string `json:"endpoint"`
				Instance            string `json:"instance"`
				Job                 string `json:"job"`
				Namespace           string `json:"namespace"`
				Pod                 string `json: "pod"`
				Prometheus          string ` json: "prometheus"`
				PrometheusReplica   string ` json: "prometheus_replica"`
				Region/*  */ string        ` json: "region"`
				Service             string ` json: "service"`
				Type                string ` json: "type"`
				Uid                 string ` json: "uid"`
			}
			Value []interface{}
		}
	}
}

/*
type ResponseExcel struct {
	ClusterInfo      QuerierData `json: "集群信息"`
	DevBusinessInfo  QuerierData `json: "开发业务信息"`
	TestBusinessInfo QuerierData `json: "测试业务信息"`
	PreBusinessInfo  QuerierData `json: "预生产业务信息"`
	ProdBusinessInfo QuerierData `json: "生产业务信息"`
	} */

type ClusterInformation struct {
	ClusterName     string  // 集群名称
	Environment     string  // 环境
	NodeNum         int     // 节点数
	TotalCPU        int     // CPU总额（core）
	TotalMemory     float64 // 内存总额（GB）
	AllocatedCPU    float64 // 已分配的CPU（核心）
	AllocatedMemory float64 // 已分配的内存（GB）
}

type BusinessInformation struct {
	Cluster          string  // 集群名称
	NameSpace        string  // 命名空间
	CPURequest       float64 // CPU请求（核心）
	MemoryRequest    float64 // 内存请求（GB）
	CPULimit         float64 // CPU上限（核心）
	MemoryLimit      float64 // 内存上限（GB）
	DeploymentCount  int     // 无状态服务数量
	StatefulsetCount int     // 有状态服务数量
	PodNum           float64 // POD数量
	PodRestartCount  string  //重启统计(重启次数)
}

// sheetName: "集群信息"
func ParseClusterInformationData(dst, env, ns, key string, jsonStr []byte) interface{} {
	clusterInformation_valueMap := make(map[string]ClusterInformation)
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

		clusterInformation_valueMap[env] = ClusterInformation{
			ClusterName:  env,
			AllocatedCPU: cpu,
		}
		WriteToExcel_ClusterInfo(dst, "集群信息", env, ns, key, startrow, clusterInformation_valueMap)

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

		clusterInformation_valueMap[env] = ClusterInformation{
			ClusterName:     env,
			AllocatedMemory: memory,
		}
		WriteToExcel_ClusterInfo(dst, "集群信息", env, ns, key, startrow, clusterInformation_valueMap)
	}
	return nil
}

func WriteToExcel_ClusterInfo(dst, sheetName, env, ns, key string, startrow int, clusterInfo_valueMap map[string]ClusterInformation) error {
	f, err := excelize.OpenFile(dst)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	info, ok := clusterInfo_valueMap[env]
	if !ok {
		return fmt.Errorf("no information for environment: %s", env)
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

// sheetName: "业务信息"
func ParseBusinessInformationData(dst, env, ns, key string, jsonStr []byte, startrow *int) interface{} {
	businessInformation_valueMap := make(map[string]BusinessInformation)
	var resp Response
	if err := json.Unmarshal(jsonStr, &resp); err != nil {
		fmt.Println("错误:", err)
		return nil
	}

	if key == "POD重启统计" {

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
				businessInformation_valueMap[env] = BusinessInformation{
					Cluster:         env,
					NameSpace:       ns,
					PodRestartCount: podRestartStr,
				}
				*startrow += 1
				fmt.Printf("===> Debug cce集群名称:%s\n    namespae:%s\n, 计算指标:%s, Pod名称:%s, 重启次数:%d, excel 表J%d行\n", env, ns, key, podRestartNameStr, podRestartCountInt, *startrow)
			}
			WriteToExcel_BusinessInformation(dst, env, ns, key, startrow, businessInformation_valueMap)

		}

	}

	/*businessInformation_valueMap[env] = ClusterInformation{
			ClusterName:  env,
			AllocatedCPU: cpu,
		}
		WriteToExcel_BusinessInformation(dst, "集群信息", env, ns, key, startrow, businessInformation_valueMap)

	} else if key == "POD数量" {
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

		businessInformation_valueMap[env] = ClusterInformation{
			ClusterName:     env,
			AllocatedMemory: memory,
		}
		WriteToExcel_BusinessInformation(dst, "集群信息", env, ns, key, startrow, businessInformation_valueMap)
	}

	if len(resp.Data.Result) == 0 || len(resp.Data.Result[0].Value) < 2 {
		fmt.Printf("utils/jsonparse.go, env: %s, ns:%s,暂无pod\n", env, ns)
		return nil
	}

	var results []string
	for _, result := range resp.Data.Result {
		if strings.Contains(result.Metric.Name, "kube_pod_container_status_restarts_total") {
			restartsStr, ok := result.Value[1].(string)
			if !ok {
				fmt.Println("Error: result.Value[1] is not a string")
				continue
			}
			restarts, err := strconv.Atoi(restartsStr)
			if err != nil {
				fmt.Println("将重新启动计数转换为整数时出错:", err)
				continue
			}

			if restarts > 0 {
				resultString := fmt.Sprintf("cce集群名称: %s,namespace: %s, Pod 名称是 %s,重启次数: %v", env, ns, result.Metric.Pod, result.Value[1])
				results = append(results, resultString)
			}

		}
	}
	if len(results) > 0 {
		return results
	} else {
		return resp.Data.Result[0].Value[1]
	} */
	return nil
}

func WriteToExcel_BusinessInformation(dst, env, ns, key string, startrow *int, businessInformation_valueMap map[string]BusinessInformation) error {
	f, err := excelize.OpenFile(dst)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	info, ok := businessInformation_valueMap[env]
	if !ok {
		return fmt.Errorf("no information for environment: %s", env)
	}

	var cell string
	var sheetName string
	switch env {
	case "hwc-sh1-dev-cluster":
		sheetName = "开发业务信息"
	case "hwc-sh1-test-cluster":
		sheetName = "测试业务信息"
	case "hwc-sh1-pre-cluster":
		sheetName = "预生产业务信息"
	case "hwc-sh1-prod-cluster":
		sheetName = "生产业务信息"
	default:
		return fmt.Errorf("unexpected environment: %s", env)
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
		cell = fmt.Sprintf("J%d", *startrow)
		f.SetCellValue(sheetName, cell, info.PodRestartCount)
		Cluster_cell := fmt.Sprintf("A%d", *startrow)
		f.SetCellValue(sheetName, Cluster_cell, info.Cluster)
		NameSpace_cell := fmt.Sprintf("B%d", *startrow)
		f.SetCellValue(sheetName, NameSpace_cell, info.NameSpace)
	default:
		return fmt.Errorf("unexpected key: %s", key)

	}

	if err = f.Save(); err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	f.Close()
	return nil
}
