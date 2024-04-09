/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calculator",
	Short: "calculate a number",
}

func Execute() {
	/* CLI Handlers */
	setupInterrupHandler()
	setupStopHandler()

	/* Root CMD */
	rootCmd.PersistentFlags().BoolP(
		"verbose",
		"v",
		false,
		"show logs and traces program",
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func setupInterrupHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		fmt.Println("\r- Wake up! Sleep has been interrupted.")
		os.Exit(0)
	}()
}

func setupStopHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTSTP)
	go func() {
		<-c
		fmt.Println("\r- Wake up! Stopped sleeping.")
		os.Exit(0)
	}()
}
