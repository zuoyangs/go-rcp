package main

import (

	"fmt"
	"github.com/spf13/viper"

	"github.com/zuoyangs/go-rcp/app/output"
	"github.com/zuoyangs/go-rcp/app/thanosapi"
	"github.com/zuoyangs/go-rcp/utils"
)

func main() {

	viper.SetConfigType("ini")
	viper.SetConfigFile("etc/config/prometheus_server.ini")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	for _, name := range viper.AllKeys() {
		fmt.Println(name)
	}

	output.Exec_output()
	thanosapi.GetClusterDetails("hwc-sh1-dev-cluster")
}
