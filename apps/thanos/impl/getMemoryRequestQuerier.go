package impl

import (
	"fmt"
	"net/http"
)

type MemoryRequestQuerier struct{}

func (m *MemoryRequestQuerier) GetResource(ClusterCategory, ClusterName, ClusterID, NameSpace string) (*http.Response, error) {
	// 查询内存请求(GB)的逻辑
	query := fmt.Sprintf("sum(kube_pod_container_resource_requests{cluster_category=\"%s\",cluster_name=\"%s\",cluster=\"%s\",namespace=\"%s\",resource=\"memory\"})/1024/1024/1024",
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
