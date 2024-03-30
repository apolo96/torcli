/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// piperCmd represents the piper command
var piperCmd = &cobra.Command{
	Use:   "piper",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("piper called")
		reader := bufio.NewReader(os.Stdin)
		s, _ := reader.ReadString('\n')
		fmt.Printf("piped in: %s\n", s)
	},
}

func init() {
	rootCmd.AddCommand(piperCmd)
}
