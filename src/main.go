package main

import "fmt"

func main() {
	filename := "config.yaml"
	config, err := ReadConfig(filename)
	if err != nil {

		fmt.Println("Error : ", err)
	}

	fmt.Println("Config :  ", config)
}
