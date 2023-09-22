package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
)

func LoadConfigValues(loadFilePath, sectionName string, keyNames []string) (map[string]string, error) {
	cfg, err := ini.Load(loadFilePath)
	if err != nil {
		fmt.Println("Error loading cluster_config.ini file:", err)
		return nil, err
	}

	section, err := cfg.GetSection(sectionName)
	if err != nil {
		fmt.Println("Error getting section:", err)
		return nil, err
	}

	configValues := make(map[string]string)
	for _, keyName := range keyNames {
		configValues[keyName] = section.Key(keyName).String()
	}

	return configValues, nil
}

func GetSectionsAndLabels(env string) (map[string]string, error) {
	v := viper.New()
	v.SetConfigType("ini")
	v.SetConfigName("config")
	v.AddConfigPath("./etc")

	// 读取配置
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("file not found.")
			return nil, err
		}
		panic(err)
	}

	labels := v.GetStringMapString(env)
	return labels, nil
}
