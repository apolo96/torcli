/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Http client to request datas a API",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{Timeout: time.Second}
		payload := &bytes.Buffer{}
		//path := "http://localhost:8080/timeout"
		path := "http://localhost:8080/request/1"
		req, err := http.NewRequest(http.MethodGet, path, payload)
		if err != nil {
			return err
		}
		res, err := client.Do(req)
		if err != nil {
			urlErr := err.(*url.Error)
			if urlErr.Timeout() {
				return fmt.Errorf("timeout %s", err)
			}
			if urlErr.Temporary() {
				return fmt.Errorf("temporary %s", err)
			}
			return fmt.Errorf("operation %s url %s error %s", urlErr.Op, urlErr.URL, urlErr.Err)
		}
		defer res.Body.Close()
		switch res.StatusCode {
		case http.StatusBadRequest:
			fmt.Printf("bad request: %v\n", res.Status)
		case http.StatusInternalServerError:
			fmt.Printf("internal service error: %v\n", res.Status)
		default:
			fmt.Printf("unexpected status code: %v\n", res.StatusCode)
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		fmt.Println(http.DetectContentType(data))
		fmt.Println(string(data))
		return nil
	},
}

var timeoutCmd = &cobra.Command{
	Use: "timeout",
	RunE: func(cmd *cobra.Command, args []string) error {
		tCmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "timeout"))
		if err := tCmd.Start(); err != nil {
			return err
		}
		errChan := make(chan error, 1)
		go func() {
			errChan <- tCmd.Wait()
		}()
		select {
		case <-time.After(time.Second * 10):
			return fmt.Errorf("timeout command")
		case err := <-errChan:
			if err != nil {
				return fmt.Errorf(err.Error())
			}
		}
		return nil
	},
}

var notfoundCmd = &cobra.Command{
	Use: "notfound",
	RunE: func(cmd *cobra.Command, args []string) error {
		eCmd := exec.Command("notfound")
		if errors.Is(eCmd.Err, exec.ErrDot) {
			fmt.Println("path lookup resolved to a local directory")
		}
		if err := eCmd.Run(); err != nil {
			if errors.Is(err, exec.ErrNotFound) {
				fmt.Println("executable failed to resolve")
			}
			return err
		}
		return nil
	},
}

var errCmd = &cobra.Command{
	Use: "error",
	RunE: func(cmd *cobra.Command, args []string) error {
		eCmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "error"))
		var out bytes.Buffer
		var stderr bytes.Buffer
		eCmd.Stdout = &out
		eCmd.Stderr = &stderr
		if err := eCmd.Run(); err != nil {
			return fmt.Errorf(stderr.String())
		}
		fmt.Println(out.String())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.AddCommand(letterCountCmd)
	execCmd.AddCommand(httpCmd)
	execCmd.AddCommand(timeoutCmd)
	execCmd.AddCommand(notfoundCmd)
	execCmd.AddCommand(errCmd)
}
