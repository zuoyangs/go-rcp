package main

import (
	"fmt"
	"time"

	thanos "gitee.com/zuoyangs/go-rcp/apps/thanos"
	"gitee.com/zuoyangs/go-rcp/utils"
)

func main() {
	dir := "./output"

	err := utils.DeleteFilesInDirectory(dir)
	if err != nil {
		fmt.Println("Error deleting files:", err)
		return
	}

	dst := fmt.Sprintf("./output/cce容器集群信息_%s.xlsx", time.Now().Format("2006-01-02_15-04-05"))
	err = utils.CopyFile("etc/cce容器集群信息.xlsx", dst)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	thanos.ThanosQuerier(dst)
}
