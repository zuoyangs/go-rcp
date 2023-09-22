package thanosapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// 封装Thanos客户端,实现查询接口调用等功能

// ConstructQuery函数用于构造Thanos查询语句
func ConstructQuery(clusterCategory, clusterName, clusterJob, clusterId string) string {
	cpuResult := queryHelper(clusterCategory, clusterName, clusterJob, "cpu", clusterId)
	memResult := queryHelper(clusterCategory, clusterName, clusterJob, "mem", clusterId)

	return cpuResult + "\n" + memResult
}

// queryHelper实现具体的指标查询逻辑,给ConstructQuery在外层调用queryHelper提供了一个简单的封装,用于获取监控采集的指标。
func queryHelper(clusterCategory, clusterName, jobType, clusterJob, clusterId string) (resultString string) {
	// ... rest of your code

	valueFloat64 := fetchAndParse(url)
	resultString = fmt.Sprintf("%.2f", valueFloat64)

	return resultString
}

func constructURL(query string) string {
	thanos_url := cfg.Section("thanos-config").Key("url").String()

	encodedQuery := url.QueryEscape(query)
	return fmt.Sprintf("%s/api/v1/query?query=%s", thanos_url, encodedQuery)
}

func fetchAndParse(urlStr string) float64 {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(urlStr)

	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	bodyByteSlice, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	var jsonResponse Response
	err = json.Unmarshal(bodyByteSlice, &jsonResponse) // Use bodyByteSlice instead of jsonStr
	if err != nil {
		fmt.Println("Error:", err)
		return 0 // Return an empty value if there's an error.
	}

	valueString, ok := jsonResponse.Data.Result[0].Value[1].(string)
	if !ok {
		fmt.Println("Error: cannot convert value to string")
		return 0
	}

	valueFloat64, err := strconv.ParseFloat(valueString, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0 // Return an empty value if there's an error.
	}

	valueFloat64 = float64(int(valueFloat64*100)) / 100

	fmt.Printf("%.2f%%\n", valueFloat64)

	return valueFloat64
}
