// Thumbnail is the API version of `nailedit`. It allows for easy and fast generation of
// image thumbnails across different platforms by selecting the fastest available generator.
package thumbnail

import (
	"io"
)

type ThumbnailOptions struct {
	// The dimensions of the output thumbnail image, in format WIDTHxHEIGHT.
	Dimensions string

	// The format of the output image, i.e. jpg, png, gif.
	Format string
}

// GenerateThumbnail generates a thumbnail of an image. It takes in an image from `input`,
// an io.Reader, and returns the generated thumbnail through `output`, an io.Writer. It tries
// multiple types of thumbnail generating tools, fastest first.
func GenerateThumbnail(input io.Reader, output io.Writer, options ThumbnailOptions) (err error) {
	// First, try to use `convert` to resize the image.
	err = resizeUsingConvert(input, output, options.Dimensions, options.Format)
	if err == nil {
		// Success, stop execution.
		return
	}

	// `convert` failed (machine may not have it), bail out to a native library.
	err = resizeUsingRez(input, output, options.Dimensions, options.Format)
	return err
}
