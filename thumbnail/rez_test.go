package thumbnail

import (
	"bytes"
	"image"
	"io/ioutil"
	"os"
	"testing"
)

func TestResizeUsingRez(t *testing.T) {
	// It should pass on errors from image.Decode.
	badImage := bytes.NewBufferString("\x00")

	err := resizeUsingRez(badImage, ioutil.Discard, "", "")
	if err != image.ErrFormat {
		t.Errorf("expected image.ErrFormat, got %v", err)
	}

	// It should error on a bad output format.
	img, err := os.Open("testdata/car.png")
	if err != nil {
		t.Fatal(err)
	}
	defer img.Close()

	err = resizeUsingRez(img, ioutil.Discard, "100x100", "badimageformat")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
