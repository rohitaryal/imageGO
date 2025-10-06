package imagego_test

import (
	"os"
	"strings"
	"testing"

	imagego "github.com/rohitaryal/imageGO/pkg/Imagego"
	prompt "github.com/rohitaryal/imageGO/pkg/Prompt"
	types "github.com/rohitaryal/imageGO/pkg/Types"
)

func TestGenerateImage(t *testing.T) {
	cookie := os.Getenv("GOOGLE_COOKIE")
	if strings.TrimSpace(cookie) == "" {
		t.Skip("Please set value of GOOGLE_COOKIE env variable.")
	}

	fx := imagego.ImageGo{Cookie: cookie}
	prompt := prompt.Prompt{
		Seed:            0,
		NumberOfImages:  3,
		Prompt:          "A cute laptop",
		AspectRatio:     types.LANDSCAPE,
		GenerationModel: types.Imagen35,
	}

	res, err := imagego.GenerateImage(&fx, prompt, false)
	if err != nil {
		t.FailNow()
	}

	if len(*res) < 1 {
		t.FailNow()
	}

	image := (*res)[0]
	path, err := image.Save("/tmp", false)
	if err != nil {
		t.FailNow()
	}

	info, err := os.Lstat(path)
	if err != nil {
		t.FailNow()
	}

	if info.Size() <= 0 {
		t.FailNow()
	}
}

func TestGetImageFromID(t *testing.T) {
	cookie := os.Getenv("GOOGLE_COOKIE")
	mediaID := os.Getenv("MEDIA_ID")
	if strings.TrimSpace(cookie) == "" {
		t.Skip("Please set value of GOOGLE_COOKIE env variable.")
	}

	if strings.TrimSpace(mediaID) == "" {
		t.Skip("Please set value of MEDIA_ID env variable.")
	}

	fx := imagego.ImageGo{Cookie: cookie}
	res, err := imagego.GetImageFromID(&fx, mediaID, true)
	if err != nil {
		t.FailNow()
	}

	if res.EncodedImage == "" {
		t.FailNow()
	}
}
