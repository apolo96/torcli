/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		/* Create command struct */
		eCmd := exec.Cmd{
			Path:   filepath.Join(os.Getenv("GOPATH"), "bin", "uppercase"),
			Args:   []string{"uppercase", "polo moroder"},
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
		/* Create OS Pipe to communicate with the child process (uppercase) */
		reader, writer, err := os.Pipe()
		if err != nil {
			return err
		}
		/* Pass a writer as file descriptor */
		eCmd.ExtraFiles = []*os.File{writer}
		if err := eCmd.Start(); err != nil {
			return err
		}
		/* Waiting for the executed command to finish */
		if err := eCmd.Wait(); err != nil {
			return err
		}
		/* Read the received data from Pipe */
		var data string
		if err := json.NewDecoder(reader).Decode(&data); err != nil {
			return err
		}
		fmt.Println(data)
		return nil
	},
}

var letterCountCmd = &cobra.Command{
	Use:   "letters",
	Short: "counter the letters in a string",
	RunE: func(cmd *cobra.Command, args []string) error {
		eCmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "lettercount"), "four")
		reader, writer, err := os.Pipe()
		if err != nil {
			return err
		}
		eCmd.ExtraFiles = []*os.File{writer}
		if err := eCmd.Run(); err != nil {
			return err
		}
		var data int
		if err := json.NewDecoder(reader).Decode(&data); err != nil {
			return err
		}
		fmt.Println(data)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.AddCommand(letterCountCmd)
}
