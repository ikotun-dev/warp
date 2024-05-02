package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port             string `yaml:"port"`
	FallbackDocument string `yaml:"fallbackDocument"`
	RootDir          string `yaml:"root"`
}

func MockConfig() *Config {
	return &Config{
		Port:             "8080",
		FallbackDocument: "index.html",
		RootDir:          "/path/to/root",
	}
}

func InitConfig() *Config {
	if _, err := os.Stat("../warp.yaml"); os.IsNotExist(err) {
		return MockConfig()
	}
	var err error
	config, err = ReadConfig("../warp.yaml")
	if err != nil {
		fmt.Println("Error : ", err)
	}
	return config
}

func ReadConfig(filename string) (*Config, error) {

	filePath := filepath.Join(".", filename)
	var config Config

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}

	log.Println("Successfully read configuration file.")

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	return &config, nil
}
