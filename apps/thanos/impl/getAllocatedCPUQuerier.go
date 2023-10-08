package impl

import (
	"fmt"
	"net/http"
)

type AllocatedCPUQuerier struct{}

func (c *AllocatedCPUQuerier) GetResource(clusterCategory, clusterName, clusterID, namespace string) (*http.Response, error) {
	// 查询已分配CPU(core)的逻辑
	query := fmt.Sprintf("sum(kube_pod_container_resource_requests{cluster_category=\"%s\",cluster_name=\"%s\",cluster=\"%s\",resource=\"cpu\"})",
		clusterCategory,
		clusterName,
		clusterID)

	resp, err := getThanosData(query)

	if err != nil {
		fmt.Println("err to get thanos data", err)
		return nil, err
	}

	return resp, nil
}
