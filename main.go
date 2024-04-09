/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"calculator/cmd"
	"calculator/cmd/logs"

	"github.com/spf13/viper"
)

func main() {
	logs.NewLogger()
	config()
	cmd.Execute()
}

func config() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		println("Warning reading in config: ", err)
	}
	viper.SetDefault("file_path", "storage/number.txt")
}
