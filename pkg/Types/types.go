// Package types contains different types and constants
package types

// Model type + constants
type Model string

const (
	Imagen3  Model = "IMAGEN_3"
	Imagen31 Model = "IMAGEN_3_1"
	Imagen35 Model = "IMAGEN_3_5"
)

// AspectRatio type + constants
type AspectRatio string

const (
	SQUARE      AspectRatio = "IMAGE_ASPECT_RATIO_SQUARE"
	PORTRAIT    AspectRatio = "IMAGE_ASPECT_RATIO_PORTRAIT"
	LANDSCAPE   AspectRatio = "IMAGE_ASPECT_RATIO_LANDSCAPE"
	UNSPECIFIED AspectRatio = "IMAGE_ASPECT_RATIO_UNSPECIFIED"
)

// DefaultHeader (mutable and shareable)
var DefaultHeader = map[string]string{
	"Origin":       "https://labs.google",
	"content-type": "application/json",
	"Referer":      "https://labs.google/fx/tools/image-fx",
}

// ImageType type + constants
type ImageType string

const (
	JPEG ImageType = "jpeg"
	JPG  ImageType = "jpg"
	JPE  ImageType = "jpe"
	PNG  ImageType = "png"
	GIF  ImageType = "gif"
	WEBP ImageType = "webp"
	SVG  ImageType = "svg"
	BMP  ImageType = "bmp"
	TIFF ImageType = "tiff"
	APNG ImageType = "apng"
	AVIF ImageType = "avif"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
}

type SessionData struct {
	User        User   `json:"user"`
	Expires     string `json:"expires"`
	AccessToken string `json:"access_token"`
}
