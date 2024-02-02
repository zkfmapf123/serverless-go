package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func GetYmlProperties[T any](ymlPath string) T {
	yamlFile, err := os.ReadFile(ymlPath)
	if err != nil {
		panic(err)
	}

	var config T
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
