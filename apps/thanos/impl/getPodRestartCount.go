package impl

import (
	"fmt"
	"net/http"
)

type PodRestartCount struct{}

func (p *PodRestartCount) GetResource(ClusterCategory, ClusterName, ClusterID, NameSpace string) (*http.Response, error) {
	// 查询POD重启数量的逻辑
	query := fmt.Sprintf("kube_pod_container_status_restarts_total{cluster_category=\"%s\",cluster_name=\"%s\",cluster=\"%s\",namespace=\"%s\"}",
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
