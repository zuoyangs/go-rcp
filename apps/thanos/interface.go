package thanos

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"

	"gitee.com/zuoyangs/go-rcp/apps/thanos/impl"
	"gitee.com/zuoyangs/go-rcp/config"
	"gitee.com/zuoyangs/go-rcp/utils/excel"
)

// 定义一个资源查询映射，键为资源类型（字符串），值为实现了ResourceQuerier接口的结构体实例
var resourceQueries = map[string]ResourceQuerier{
	"已分配CPU(core)": &impl.AllocatedCPUQuerier{},    // CPU 已分配查询
	"已分配内存(GB)":    &impl.AllocatedMemoryQuerier{}, // 内存 已分配查询
	"POD重启统计":      &impl.PodRestartCount{},        // Pod重启次数

	/* 	"CPU请求(core)":  &impl.CPURequestQuerier{},      // CPU requests查询
		"CPU上限(core)":  &impl.CPULimitQuerier{},        // CPU limits查询
	  	"内存请求(GB)":     &impl.MemoryRequestQuerier{},   // 内存 requests查询
		"内存上限(GB)":     &impl.MemLimitQuerier{},        // 内存 limits查询
		"无状态服务数量":      &impl.K8sDeploymentCount{},     // k8s 无状态服务数量
		"有状态服务数量":      &impl.K8sStatefulSetCount{},    // k8s 有状态服务数量
		"POD数量":        &impl.PodNum{},                 // Pod数量*/
}

// 定义 ResourceQuerier 模块接口
type ResourceQuerier interface {
	// 查询 thanos resource 接口
	GetResource(clusterCategory, clusterName, clusterID, namespace string) (*http.Response, error)
}

func ThanosQuerier(dst string) {
	envs := []string{"hwc-sh1-prod-cluster", "hwc-sh1-pub-cluster", "hwc-sh1-pre-cluster", "hwc-sh1-test-cluster", "hwc-sh1-dev-cluster"}
	//envs := []string{"hwc-sh1-dev-cluster"}
	config.Init()
	excel.InsertCurrentTimeToExcel(dst, "集群信息")
	for _, env := range envs {
		_, err := config.GetSectionsAndLabels(env)
		if err != nil {
			fmt.Println("Error getting labels: ", err)
			continue
		}

		ClusterCategory, _ := config.GetKey(env, "cluster_category")
		ClusterName, _ := config.GetKey(env, "cluster_name")
		ClusterID, _ := config.GetKey(env, "cluster")

		var startrow int = 1

		nsfile, err := os.Open("etc/ns.txt")
		if err != nil {
			fmt.Println("无法打开文件:", err)
			return
		}
		defer nsfile.Close()

		scanner := bufio.NewScanner(nsfile)
		for scanner.Scan() {
			ns := scanner.Text()
			for key := range resourceQueries {
				respString, err := resourceQueries[key].GetResource(ClusterCategory,
					ClusterName,
					ClusterID,
					ns)
				if err != nil {
					fmt.Printf("Error! 业务信息表获取 %s 失败: %v\n", key, err)
				} else {

					body, err := io.ReadAll(respString.Body)
					if err != nil {
					}
					impl.ParseClusterInfoData(dst, env, ns, key, body)
					impl.ParseBusinessInfoData(dst, env, ns, key, body, &startrow)
				}
			}
		}
	}
}
