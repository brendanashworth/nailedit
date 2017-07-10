package thumbnail

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateThumbnail(t *testing.T) {
	// Generate a thumbnail for testdata/car.png and check the md5 sum
	// with different options. CWD = /thumbnail/
	img1, err := os.Open("testdata/car.png")
	if err != nil {
		t.Fatal(err)
	}
	defer img1.Close()

	options := ThumbnailOptions{
		Dimensions:     "128x128",
		Format:         "jpeg",
		AvoidLibraries: AvoidRez,
	}

	var expectedDefault string = "b89a02391af14cf00d285fd278eb66a6"
	hash := md5.New()

	err = GenerateThumbnail(img1, hash, options)
	if err != nil && err != ErrNoLibrary {
		t.Error(err)
	} else {
		if expectedDefault != fmt.Sprintf("%x", hash.Sum(nil)) {
			t.Error("default options hash does not match expected")
		}
	}

	// Now do it for a jpg.
	img2, err := os.Open("testdata/beach.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer img2.Close()

	options = ThumbnailOptions{
		Dimensions:     "64x64",
		Format:         "gif",
		AvoidLibraries: AvoidConvert,
	}

	var expectedSmallGIF string = "5da7b19111cce28aa31142d8344395bc"
	hash = md5.New()

	err = GenerateThumbnail(img2, hash, options)
	if err != nil {
		t.Error(err)
	} else {
		if expectedSmallGIF != fmt.Sprintf("%x", hash.Sum(nil)) {
			t.Error("changed options hash does not match expected")
		}
	}

	// This should fail with no libraries.
	options = ThumbnailOptions{
		Dimensions:     "1x1",
		Format:         "jpeg",
		AvoidLibraries: AvoidConvert | AvoidRez,
	}

	err = GenerateThumbnail(img2, ioutil.Discard, options)
	if err != ErrNoLibrary {
		t.Errorf("expected ErrNoLibrary, got %s", err.Error())
	}
}
