package impl

import (
	"errors"
	"fmt"
	"net/http"

	"gitee.com/zuoyangs/go-rcp/config"
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

func getThanosData(encodedQuery string) (*http.Response, error) {

	config.Init()
	thanos_url, err := config.GetKey("thanos-config", "url")
	if err != nil {
		return nil, fmt.Errorf("Failed to get thanos URL: %v", err)
	}

	url := fmt.Sprintf("%s/api/v1/query?query=%s", thanos_url, encodedQuery)

	//fmt.Printf("\nthanos query url(apps/thanos/impl/impl.go): %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	if resp.Body == nil {
		return nil, errors.New("response body is empty")
	}
	return resp, nil
}
