package impl

import (
	"fmt"
	"net/http"
)

type K8sDeploymentCount struct{}

func (k *K8sDeploymentCount) GetResource(ClusterCategory, ClusterName, ClusterID, NameSpace string) (*http.Response, error) {
	// 查询无状态服务数量的逻辑
	query := fmt.Sprintf("count(kube_deployment_created{cluster_category=\"%s\",cluster_name=\"%s\",cluster=\"%s\",namespace=\"%s\"})",
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
