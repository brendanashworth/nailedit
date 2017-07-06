// Thumbnail is the API version of `nailedit`. It allows for easy and fast generation of
// image thumbnails across different platforms by selecting the fastest available generator.
package thumbnail

import (
	"errors"
	"io"
)

const (
	AvoidConvert = 1 << iota
	AvoidRez
)

// No library was available to convert the image with. They may all be disabled
// or unreachable or were causing errors. This may be a problem with the image
// itself.
var ErrNoLibrary = errors.New("no library available")

type ThumbnailOptions struct {
	// The dimensions of the output thumbnail image, in format WIDTHxHEIGHT.
	Dimensions string

	// The format of the output image, i.e. jpg, png, gif.
	Format string

	// The collection of libraries to not test, i.e. AvoidConvert to avoid ImageMagick's `convert`.
	AvoidLibraries int
}

// GenerateThumbnail generates a thumbnail of an image. It takes in an image from `input`,
// an io.Reader, and returns the generated thumbnail through `output`, an io.Writer. It tries
// multiple types of thumbnail generating tools, fastest first.
func GenerateThumbnail(input io.Reader, output io.Writer, options ThumbnailOptions) (err error) {
	// First, try to use `convert` to resize the image.
	if options.AvoidLibraries&AvoidConvert == 0 {
		err = resizeUsingConvert(input, output, options.Dimensions, options.Format)
		if err == nil {
			// Success, stop execution.
			return
		}
	}

	// Next try to use rez, a native Go library.
	if options.AvoidLibraries&AvoidRez == 0 {
		err = resizeUsingRez(input, output, options.Dimensions, options.Format)
		return err
	}

	return ErrNoLibrary
}
