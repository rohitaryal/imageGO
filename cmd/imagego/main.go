package main

import (
	"flag"
	"fmt"
	"strings"

	imagego "github.com/rohitaryal/imageGO/pkg/Imagego"
	prompt "github.com/rohitaryal/imageGO/pkg/Prompt"
	types "github.com/rohitaryal/imageGO/pkg/Types"
)

func main() {
	generate := flag.Bool("image", false, "Generate image from a prompt")
	promptStr := flag.String("prompt", "Purple cat", "Textual description for the image")
	seed := flag.Int("seed", 0, "A specific number that serves as the starting point")
	count := flag.Int("count", 1, "Number of images to generate")
	aspectRatio := flag.String("size", "LANDSCAPE", "Aspect ratio of the image")
	model := flag.String("model", "IMAGEN35", "Model to use for generation")
	dir := flag.String("dir", ".", "Destination directory to save images")

	caption := flag.Bool("caption", false, "Generate description from an image")

	fetch := flag.Bool("fetch", false, "Fetch generated images using unique media ID")
	mediaID := flag.String("id", "", "Unique media generation id")

	cookie := flag.String("cookie", "", "User account cookie")
	verbose := flag.Bool("verbose", false, "Extra logs")

	flag.Usage = func() {
		fmt.Println("imagego [flags] --cookie [cookie]")
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Available sizes: SQUARE, LANDSCAPE, PORTRAIT, UNSPECIFIED")
		fmt.Println("Available models: IMAGEN3, IMAGEN31, IMAGEN35")
	}

	flag.Parse()

	if strings.TrimSpace(*cookie) == "" {
		fmt.Println("Cookie value is missing")
		return
	}
	a := imagego.ImageGo{Cookie: *cookie}

	if *generate {
		*model = strings.ToLower(*model)
		switch *model {
		case "imagen3":
			*model = string(types.Imagen3)
		case "imagen31":
			*model = string(types.Imagen31)
		case "imagen35":
			*model = string(types.Imagen35)
		default:
			fmt.Println("Unknown model: ", *model)
			return
		}

		*aspectRatio = strings.ToLower(*aspectRatio)
		switch *aspectRatio {
		case "square":
			*aspectRatio = string(types.SQUARE)
		case "landscape":
			*aspectRatio = string(types.LANDSCAPE)
		case "portrait":
			*aspectRatio = string(types.PORTRAIT)
		case "unspecified":
			*aspectRatio = string(types.UNSPECIFIED)
		default:
			fmt.Println("Unknown aspect ratio: ", *aspectRatio)
			return
		}

		p := prompt.Prompt{
			Seed:            *seed,
			Prompt:          *promptStr,
			NumberOfImages:  *count,
			AspectRatio:     types.AspectRatio(*aspectRatio),
			GenerationModel: types.Model(*model),
		}

		res, err := imagego.GenerateImage(&a, p, *verbose)
		if err != nil {
			fmt.Println("Failed to generate iamge.")
			fmt.Println(err)
		}

		for _, value := range *res {
			path, err := value.Save(*dir, *verbose)
			if err != nil {
				fmt.Println("Failed to save an image")
			} else {
				fmt.Println("[+] Saved image: " + path)
			}
		}
	} else if *caption {
	} else if *fetch {
		if strings.TrimSpace(*mediaID) == "" {
			fmt.Println("Media ID is missing.")
			return
		}

		res, err := imagego.GetImageFromID(&a, *mediaID, *verbose)
		if err != nil {
			fmt.Println("Failed to fetch image")
			fmt.Println(err)
			return
		}

		path, err := res.Save(*dir, *verbose)
		if err != nil {
			fmt.Println("Failed to save fetched image")
			fmt.Println(err)
			return
		}
		fmt.Println("[+] Image saved: " + path)
	} else {
		fmt.Println("Please see the usage in README.md: https://github.com/rohitaryal/imageGO")
	}
}
