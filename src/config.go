package main

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	port             int    `yaml:"port"`
	fallbackDocument string `yaml:"fallbackDocument"`
	rootDir          string `yaml:"root"`
}
