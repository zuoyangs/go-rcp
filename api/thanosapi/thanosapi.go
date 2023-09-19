package thanosapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"gopkg.in/ini.v1"
)

func GetClusterDetails(ssection string) string {
	cfg, err := ini.Load("etc/cluster_config.ini")
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

	return ConstructQuery(clusterCategory, clusterName, clusterJob, clusterMode, clusterId)
}

func ConstructQuery(clusterCategory, clusterName, clusterJob, clusterMode, clusterId string) string {

	query :=
		fmt.Sprintf(
			"avg(sum by (instance)(irate(node_cpu_seconds_total{cluster_category=\"%s\",cluster_name=\"%s\",clusterJob=\"%s\",clusterMode=\"%s\",clusterId=\"%s\"}[7d])) * 100)",
			clusterCategory,
			clusterName,
			clusterJob,
			clusterMode,
			clusterId,
		)

	encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("https://thanos.wehotelio.com/api/v1/query?query=%s", encodedQuery)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
		return ""
	}

	return string(body)
}
