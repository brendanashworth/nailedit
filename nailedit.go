package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brendanashworth/nailedit/thumbnail"
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

	options := thumbnail.ThumbnailOptions{
		Dimensions: *dimensionFlag,
		Format:     *formatFlag,
	}

	err := thumbnail.GenerateThumbnail(os.Stdin, os.Stdout, options)
	if err != nil {
		fmt.Errorf("error occurred generating thumbnail: %s\n", err.Error())
		os.Exit(1)
	}
}
