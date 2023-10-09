package main

import (
	"fmt"
	"os"
	"time"

	thanos "gitee.com/zuoyangs/go-rcp/apps/thanos"
	"gitee.com/zuoyangs/go-rcp/utils"
)

func main() {
	dir := "./output"

	// 检查目录是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("创建目录失败:", err)
			return
		}
		fmt.Println("目录创建成功:", dir)
	}

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
