package thumbnail

import (
	"io"
	"os/exec"
)

// Attempts to resize an image using the `convert` command line tool.
func resizeUsingConvert(input io.Reader, output io.Writer, dimensions, format string) error {
	// Call into convert.
	cmd := exec.Command("convert", "-", "-thumbnail", dimensions, format+":-")
	cmd.Stdin = input
	cmd.Stdout = output

	err := cmd.Run()
	return err
}
