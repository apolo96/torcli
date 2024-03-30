/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"calculator/storage"
	"fmt"

	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		storage.SaveNumber(0)
		fmt.Println(0)
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
