/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"calculator/cmd/validator"
	"calculator/storage"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// multiplyCmd represents the multiply command
var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if err := validator.ValidateArgs(args); err != nil {
			fmt.Println(err)
		}
		number, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Printf("unable to parse input[%s]: %v", args[0], err)
			return
		}
		value := storage.GetNumber() * number
		storage.SaveNumber(value)
		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
}
