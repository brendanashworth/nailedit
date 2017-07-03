package main

import (
	"flag"
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
	if err != nil {
		// `convert` failed (machine may not have it), use resize, a Go library.

		os.Exit(1)
		// execution stops
	}

	// Output is piped straight to STDOUT, we're done!
}
