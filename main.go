package main

import (
	"fmt"

	execl "github.com/zuoyangs/go-rcp/app/excel"
	"github.com/zuoyangs/go-rcp/app/thanos"
	"github.com/zuoyangs/go-rcp/utils"
)

func main() {
	envs := []string{"hwc-sh1-dev-cluster", "hwc-sh1-test-cluster", "hwc-sh1-pre-cluster", "hwc-sh1-prod-cluster"}

	for _, env := range envs {
		labels, err := utils.GetSectionsAndLabels(env)
		fmt.Println(env, labels)
		if err != nil {
			fmt.Println("Error getting labels: ", err)
			continue
		}

		avgCPUUsage, avgMemUsage, peakCPUUsage, peakMemUsage, err := thanos.GetClusterDetails(env, labels)

		if err != nil {
			fmt.Println("Error getting cluster details: ", err)
			continue
		}

		if err = execl.WriteToExcel(avgCPUUsage, avgMemUsage, peakCPUUsage, peakMemUsage); err != nil {
			fmt.Println("Error writing to excel: ", err)
		}
	}
}
