package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port             string `yaml:"port"`
	FallbackDocument string `yaml:"fallbackDocument"`
	RootDir          string `yaml:"root"`
}

func InitConfig() *Config {
	config, err := ReadConfig("config.yaml")
	if err != nil {
		fmt.Println("Error : ", err)
	}
	return config
}

func ReadConfig(filename string) (*Config, error) {

	filePath := filepath.Join("./examples/static", filename)
	var config Config

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}

	fmt.Println("YAML file contents:", string(yamlFile))

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	return &config, nil
}
