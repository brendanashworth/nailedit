package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/bamiaux/rez"
)

// Attempts to resize an image using the rez library.
func resizeUsingRez(inputFile, outputFile *os.File, dimensions, format string) error {
	// Get images
	input, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	var width int
	var height int
	fmt.Sscanf(dimensions, "%dx%d", &width, &height)

	typedImage, ok := input.(*image.YCbCr)
	if !ok {
		return fmt.Errorf("input picture must be YCbCr")
	}

	var output image.Image
	output = image.NewYCbCr(image.Rect(0, 0, width, height), typedImage.SubsampleRatio)

	// Pump it into the library.
	err = rez.Convert(output, input, rez.NewBilinearFilter())
	if err != nil {
		return err
	}

	// Write the output image to the output file.
	if format == "jpeg" {
		err = jpeg.Encode(outputFile, output, nil)
	} else if format == "png" {
		err = png.Encode(outputFile, output)
	} else if format == "gif" {
		err = gif.Encode(outputFile, output, nil)
	} else {
		err = fmt.Errorf("unsupported image format '%s' for rez (use jpeg, png, or gif)", format)
	}

	return err
}
