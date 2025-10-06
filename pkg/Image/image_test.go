package image_test

import (
	"os"
	"testing"

	image "github.com/rohitaryal/imageGO/pkg/Image"
)

func TestSave(t *testing.T) {
	image := &image.Image{
		// Smalllest base64 image
		EncodedImage: "iVBORw0KGgoAAAANSUhEUgAAAAgAAAAIAQMAAAD+wSzIAAAABlBMVEX///+/v7+jQ3Y5AAAADklEQVQI12P4AIX8EAgALgAD/aNpbtEAAAAASUVORK5CYII=",
	}

	path, err := image.Save(".", true)
	if err != nil {
		t.FailNow()
	}

	// Remove it too
	os.Remove(path)
}
