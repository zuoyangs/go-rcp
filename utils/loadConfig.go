package utils

import (
	"fmt"

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
