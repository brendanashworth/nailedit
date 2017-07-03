package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Set up the command line flags.
	dimensionFlag := flag.String("dimensions", "128x128", "dimensions of the output thumbnail image")
	formatFlag := flag.String("format", "jpeg", "format of the output thumbnail image")
	helpFlag := flag.Bool("help", false, "print help")

	flag.Parse()

	if *helpFlag {
		fmt.Printf("\tnailedit\n\n")
		flag.PrintDefaults()
		return
	}

	// It works directly with stdin and stdout.
	inputFile := os.Stdin
	outputFile := os.Stdout

	// Call into convert.
	cmd := exec.Command("convert", "-", "-resize", *dimensionFlag, (*formatFlag)+":-")
	cmd.Stdin = inputFile
	cmd.Stdout = outputFile

	err := cmd.Run()
	if err != nil {
		fmt.Errorf("convert failed, %s\n", err.Error())
		os.Exit(1)
		// execution ends
	}

	// Output from the `convert` command is piped straight to STDOUT for handling
	// by the user.
}
