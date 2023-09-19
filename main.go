package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("etc/cce容器集群信息.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows("集群信息")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, cell := range row {
			fmt.Print(cell, "\t")
		}
		fmt.Println()
	}

}
