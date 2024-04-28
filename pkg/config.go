package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port             int    `yaml:"port"`
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

	var config Config

	yamlFile, err := os.ReadFile("config.yaml")
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
