package thumbnail

import (
	"crypto/md5"
	"fmt"
	"os"
	"testing"
)

func TestGenerateThumbnail(t *testing.T) {
	// Generate a thumbnail for testdata/beach.jpg and check the md5 sum
	// with different options. CWD = /thumbnail/
	img1, err := os.Open("testdata/beach.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer img1.Close()

	options := ThumbnailOptions{
		Dimensions: "128x128",
		Format:     "jpeg",
	}

	var expectedDefault string = "3a696fcddda2a8e089aa2296070840cd"
	hash := md5.New()

	err = GenerateThumbnail(img1, hash, options)
	if err != nil {
		t.Error(err)
	} else {
		if expectedDefault != fmt.Sprintf("%x", hash.Sum(nil)) {
			t.Error("default options hash does not match expected")
		}
	}

	// Now do it for a gif.
	img2, err := os.Open("testdata/car.gif")
	if err != nil {
		t.Fatal(err)
	}
	defer img2.Close()

	options = ThumbnailOptions{
		Dimensions: "64x64",
		Format:     "gif",
	}

	var expectedSmallGIF string = "4b2a4809ae5d62005ecd49260ac8a3b9"
	hash = md5.New()

	err = GenerateThumbnail(img2, hash, options)
	if err != nil {
		t.Error(err)
	} else {
		if expectedSmallGIF != fmt.Sprintf("%x", hash.Sum(nil)) {
			t.Error("changed options hash does not match expected")
		}
	}
}
