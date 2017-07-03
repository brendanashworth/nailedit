package main

import (
	"os"
	"os/exec"
)

// Attempts to resize an image using the `convert` command line tool.
func resizeUsingConvert(input, output *os.File, dimensions, format string) error {
	// Call into convert.
	cmd := exec.Command("convert", "-", "-resize", dimensions, format+":-")
	cmd.Stdin = input
	cmd.Stdout = output

	err := cmd.Run()
	return err
}
