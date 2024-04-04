/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

// systemCmd represents the system command
var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("system called")
	},
}

var envCmd = &cobra.Command{
	Use: "env",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println("working directory ", dir)
		if err := os.Setenv("WORKING_DIR", dir); err != nil {
			return err
		}
		fmt.Println(os.ExpandEnv("WORKING_DIR=${WORKING_DIR}"))
		if err := os.Unsetenv("WORKING_DIR"); err != nil {
			return nil
		}
		fmt.Println(os.ExpandEnv("WORKING_DIR=${WORKING_DIR}"))
		fmt.Println("There are env vars: ", len(os.Environ()))
		for _, env := range os.Environ() {
			fmt.Println(env)
		}
		return nil
	},
}

var timeCmd = &cobra.Command{
	Use: "time",
	RunE: func(cmd *cobra.Command, args []string) error {
		start := time.Now()
		fmt.Println("start time ", start)
		time.Sleep(1 * time.Second)
		elapsed := time.Until(start)
		fmt.Println("elapsed time ", elapsed)
		return nil
	},
}

var runtimeCmd = &cobra.Command{
	Use: "runtime",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Sistema operativo:", runtime.GOOS)
		fmt.Println("Arquitectura:", runtime.GOARCH)
		fmt.Println("Go Root:", runtime.GOROOT())
		fmt.Println("Compilador:", runtime.Compiler)
		fmt.Println("Núm. de CPU:", runtime.NumCPU())
		fmt.Println("Núm. de Goroutines:", runtime.NumGoroutine())
		fmt.Println("Versión:", runtime.Version())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(systemCmd)
	systemCmd.AddCommand(envCmd)
	systemCmd.AddCommand(timeCmd)
	systemCmd.AddCommand(runtimeCmd)
}
