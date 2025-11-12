package main

import (
	"fmt"

	"github.com/kumose/xcmd"
)

func main() {
	// Create Cmd, buffered output
	envCmd := xcmd.NewCmd("env")

	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()

	// Print each line of STDOUT from Cmd
	for _, line := range status.Stdout {
		fmt.Println(line)
	}
}
