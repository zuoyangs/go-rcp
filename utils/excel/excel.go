package excel

/*
import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Result struct {
	Metric map[string]string `json:"metric"`
	Value  []interface{}     `json:"value"`
}

type Data struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}

type Response struct {
	Data struct {
		Result []struct {
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func ProcessData(bodyBytes []byte) (string, error) {
	var resp Response
	if err := json.Unmarshal(bodyBytes, &resp); err != nil {
		return "", fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	valueStr, ok := resp.Data.Result[0].Value[1].(string)
	if !ok {
		return "", fmt.Errorf("cannot convert to string")
	}

	valueFloat, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return "", fmt.Errorf("cannot convert to float64: %v", err)
	}

	valueRounded := math.Round(valueFloat*100) / 100
	valueString := fmt.Sprintf("%.1f", valueRounded)

	return valueString, nil
}

func writeToExcel(filePath, sheetName, cell, valueString string) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	f.SetCellValue(sheetName, cell, valueString)

	if err = f.Save(); err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}
	f.Close()
	return nil
}

type QuerierData struct {
	AllocatedCPUQuerier    interface{} `json:"已分配CPU查询"`
	CPURequestQuerier      interface{} `json:"CPU请求查询"`
	CPULimitQuerier        interface{} `json:"CPU限制查询"`
	AllocatedMemoryQuerier interface{} `json:"已分配内存查询"`
	MemoryRequestQuerier   interface{} `json:"内存请求查询"`
	MemLimitQuerier        interface{} `json:"内存限制查询"`
	K8sDeploymentCount     interface{} `json:"k8s无状态服务数量"`
	K8sStatefulSetCount    interface{} `json:"k8s有状态服务数量"`
	PodNum                 interface{} `json:"Pod数量"`
	PodRestartCount        interface{} `json:"Pod重启次数"`
}

type ResponseExcel struct {
	ClusterInfo      QuerierData `json: "集群信息"`
	DevBusinessInfo  QuerierData `json: "开发业务信息"`
	TestBusinessInfo QuerierData `json: "测试业务信息"`
	PreBusinessInfo  QuerierData `json: "预生产业务信息"`
	ProdBusinessInfo QuerierData `json: "生产业务信息"`
}

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
	Namespace         string  // 命名空间
	CPURequest        float64 // CPU请求（核心）
	MemoryRequest     float64 // 内存请求（GB）
	CPULimit          float64 // CPU上限（核心）
	MemoryLimit       float64 // 内存上限（GB）
	StatelessServices string  // 无状态服务
	StatefulServices  string  // 有状态服务
	PODNumbers        int     // POD数量
	RestartStatistics int     //重启统计(重启次数)
}
*/
