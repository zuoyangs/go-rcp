package impl

import (
	"fmt"
	"net/http"
)

type PodNum struct{}

func (p *PodNum) GetResource(ClusterCategory, ClusterName, ClusterID, NameSpace string) (*http.Response, error) {
	// 查询POD数量的逻辑
	query := fmt.Sprintf("count(kube_pod_info{cluster_category=\"%s\",cluster_name=\"%s\",cluster=\"%s\",namespace=\"%s\"})",
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
