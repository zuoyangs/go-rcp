package config

// 处理配置文件,如加载ini文件等

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *viper.Viper

func Init() {

	cfg = viper.New()
	cfg.SetConfigType("ini")
	cfg.SetConfigName("config")
	cfg.AddConfigPath("./etc")

	if err := cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("file not found")
			return
		}
		panic(err)
	}
}

func GetSectionsAndLabels(env string) (map[string]string, error) {
	labels := cfg.GetStringMapString(env)
	return labels, nil
}

func GetKey(section, key string) (string, error) {

	value := cfg.GetString(fmt.Sprintf("%s.%s", section, key))

	return string(value), nil
}
