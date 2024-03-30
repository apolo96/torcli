/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// sleepCmd represents the sleep command
var sleepCmd = &cobra.Command{
	Use:   "sleep",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sleep called")
		for {
			fmt.Println("ZzZzz")
			time.Sleep(time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(sleepCmd)
}
