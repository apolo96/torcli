/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("convert called")
	},
}

var moneySubCmd = &cobra.Command{
	Use:   "money",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("money called")
		from, _ := cmd.Flags().GetString("from")
		fmt.Printf("From: %s\n", from)
		to, _ := cmd.Flags().GetString("to")
		fmt.Printf("To: %s\n", to)
		return nil
	},
}

var lengthSubCmd = &cobra.Command{
	Use:   "length",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		from, _ := cmd.Flags().GetString("from")
		fmt.Printf("From: %s\n", from)
		to, _ := cmd.Flags().GetString("to")
		fmt.Printf("To: %s\n", to)
		return nil
	},
}

func init() {
	/* Convert Money with flags */
	convertCmd.AddCommand(moneySubCmd)
	convertCmd.PersistentFlags().String("from", "", "Current unit")
	convertCmd.MarkPersistentFlagRequired("from")
	convertCmd.PersistentFlags().String("to", "", "Unit to convert")
	convertCmd.MarkPersistentFlagRequired("to")
	/* Convert Length with args */
	convertCmd.AddCommand(lengthSubCmd)
	/* Convert */
	rootCmd.AddCommand(convertCmd)
}
