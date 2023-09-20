package thanosapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"gopkg.in/ini.v1"
)

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}

type Result struct {
	Metric Metric        `json:"metric"`
	Value  []interface{} `json:"value"`
}

type Metric struct{}

type Value json.Number

func GetClusterDetails1(ssection string) string {
	cfg, err := ini.Load("etc/config/prometheus_server.ini")
	if err != nil {
		fmt.Println("Error loading cluster_config.ini file:", err)
		return ""
	}

	section, err := cfg.GetSection(ssection)
	if err != nil {
		fmt.Println("Error getting section:", err)
		return ""
	}

	clusterCategory := section.Key("cluster_category").String()
	clusterName := section.Key("cluster_name").String()
	clusterJob := section.Key("job").String()
	clusterMode := section.Key("mode").String()
	clusterId := section.Key("cluster").String()

	cpuResult := CPU_ConstructQuery(clusterCategory, clusterName, clusterJob, clusterMode, clusterId)
	memResult := MEM_ConstructQuery(clusterCategory, clusterName, clusterJob, clusterId)
	return cpuResult + "\n" + memResult
}

func CPU_ConstructQuery(clusterCategory, clusterName, clusterJob, clusterMode, clusterId string) string {

	cfg, err := ini.Load("etc/config/thanos_server.ini")
	if err != nil {
		fmt.Println("Error loading cluster_config.ini file:", err)
		return ""
	}
	section, err := cfg.GetSection("thanos-config")
	if err != nil {
		fmt.Println("Error getting section:", err)
		return ""
	}
	thanos_url := section.Key("url").String()

	query :=
		fmt.Sprintf(
			"avg(sum by (instance)(irate(node_cpu_seconds_total{cluster_category=\"%s\",cluster_name=\"%s\",job=\"%s\",mode=\"%s\",cluster=\"%s\"}[7d])) * 100)",
			clusterCategory,
			clusterName,
			clusterJob,
			clusterMode,
			clusterId,
		)

	encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("%s/api/v1/query?query=%s", thanos_url, encodedQuery)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
		return ""
	}
	defer resp.Body.Close()

	bodyByteSlice, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
		return ""
	}

	var jsonResponse Response
	err = json.Unmarshal(bodyByteSlice, &jsonResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	valueString, ok := jsonResponse.Data.Result[0].Value[1].(string)
	if !ok {
		fmt.Println("Error: cannot convert value to string")
		return ""
	}

	valueFloat64, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	valueFloat64 = float64(int(valueFloat64*100)) / 100

	fmt.Printf("%.2f%%\n", valueFloat64)

	return fmt.Sprintf("%.2f", valueFloat64)
}

func MEM_ConstructQuery(clusterCategory, clusterName, clusterJob, clusterId string) string {

	cfg, err := ini.Load("etc/config/thanos_server.ini")
	if err != nil {
		fmt.Println("Error loading cluster_config.ini file:", err)
		return ""
	}
	section, err := cfg.GetSection("thanos-config")
	if err != nil {
		fmt.Println("Error getting section:", err)
		return ""
	}
	thanos_url := section.Key("url").String()

	query :=
		fmt.Sprintf("avg((node_memory_MemTotal_bytes{cluster_category=\"%s\",cluster_name=\"%s\",job=\"%s\",cluster=\"%s\"} - node_memory_MemFree_bytes{cluster_category=\"%s\",cluster_name=\"%s\",job=\"%s\",cluster=\"%s\"}) / (node_memory_MemTotal_bytes{cluster_category=\"%s\",cluster_name=\"%s\",job=\"%s\",cluster=\"%s\"} )) * 100",
			clusterCategory,
			clusterName,
			clusterJob,
			clusterId,
			clusterCategory,
			clusterName,
			clusterJob,
			clusterId,
			clusterCategory,
			clusterName,
			clusterJob,
			clusterId,
		)

	encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("%s/api/v1/query?query=%s", thanos_url, encodedQuery)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
		return ""
	}
	defer resp.Body.Close()

	bodyByteSlice, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
		return ""
	}

	var jsonResponse Response
	err = json.Unmarshal(bodyByteSlice, &jsonResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	valueString, ok := jsonResponse.Data.Result[0].Value[1].(string)
	if !ok {
		fmt.Println("Error: cannot convert value to string")
		return ""
	}

	valueFloat64, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	valueFloat64 = float64(int(valueFloat64*100)) / 100

	fmt.Printf("%.2f%%", valueFloat64)

	return fmt.Sprintf("%.2f", valueFloat64)
}

func GetClusterDetails(env string, labels map[string]string) (avgCPUUsage float64, avgMemUsage float64, peakCPUUsage float64, peakMemUsage float64, err error) {
	return
}
