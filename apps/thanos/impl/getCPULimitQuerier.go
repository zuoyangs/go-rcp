package impl

import (
	"fmt"
	"net/http"
)

type CPULimitQuerier struct{}

func (c *CPULimitQuerier) GetResource(ClusterCategory, ClusterName, ClusterID, NameSpace string) (*http.Response, error) {
	// 查询CPU上限(core)的逻辑
	query := fmt.Sprintf("sum(kube_pod_container_resource_limits{cluster_category=\"%s\",cluster_name=\"%s\",cluster=\"%s\",namespace=\"%s\",resource=\"cpu\"})",
		ClusterCategory,
		ClusterName,
		ClusterID,
		NameSpace)

	resp, err := getThanosData(query)

	if err != nil {
		fmt.Println("err to get thanos data", err)
		return nil, err
	}

	return resp, nil
}
