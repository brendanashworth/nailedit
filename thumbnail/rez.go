package thumbnail

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/bamiaux/rez"
)

// Attempts to resize an image using the rez library.
func resizeUsingRez(inputReader io.Reader, outputWriter io.Writer, dimensions, format string) error {
	// Get images
	input, _, err := image.Decode(inputReader)
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
		err = jpeg.Encode(outputWriter, output, nil)
	} else if format == "png" {
		err = png.Encode(outputWriter, output)
	} else if format == "gif" {
		err = gif.Encode(outputWriter, output, nil)
	} else {
		err = fmt.Errorf("unsupported image format '%s' for rez (use jpeg, png, or gif)", format)
	}

	return err
}
