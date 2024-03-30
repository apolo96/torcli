/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"calculator/cmd"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Warning reading in config: ", err)
	}
	viper.SetDefault("file_path", "storage/number.txt")
	cmd.Execute()
}
