package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Set up the command line flags.
	dimensionFlag := flag.String("dimensions", "128x128", "dimensions of the output thumbnail image")
	formatFlag := flag.String("format", "jpeg", "format of the output thumbnail image")
	helpFlag := flag.Bool("help", false, "print help")

	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		return
	}

	// It works directly with stdin and stdout.
	inputFile := os.Stdin
	outputFile := os.Stdout

	// First, try to use `convert` to resize the image.
	err := resizeUsingConvert(inputFile, outputFile, *dimensionFlag, *formatFlag)
	if err == nil {
		// Success, stop execution.
		return
	}

	// `convert` failed (machine may not have it), bail out to a native library.
	err = resizeUsingRez(inputFile, outputFile, *dimensionFlag, *formatFlag)
	if err != nil {
		fmt.Printf("error occurred using rez: %s\n", err.Error())

		// rez failed, bail the thumbnail entirely.
		os.Exit(1)
	}

}
