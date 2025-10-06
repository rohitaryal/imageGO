// Package imagego is the main package
package imagego

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	utils "github.com/rohitaryal/imageGO/internal/Utils"
	account "github.com/rohitaryal/imageGO/pkg/Account"
	image "github.com/rohitaryal/imageGO/pkg/Image"
	prompt "github.com/rohitaryal/imageGO/pkg/Prompt"
	types "github.com/rohitaryal/imageGO/pkg/Types"
)

type ImageGo struct {
	Cookie  string
	account account.Account
}

func GenerateImage(i *ImageGo, p prompt.Prompt, verbose bool) (*[]image.Image, error) {
	if strings.TrimSpace(i.Cookie) == "" {
		return nil, fmt.Errorf("user cookie is required")
	}

	// If Parent struct cookie doesn't match with the account inside it
	// Then probably user tried to change the cookie, so lets make this sync with
	// account too.
	if i.Cookie != i.account.Cookie {
		if verbose {
			fmt.Println("Cookie don't match, lets refresh")
		}
		i.account.Cookie = i.Cookie
		err := i.account.RefreshSession(verbose)
		if err != nil {
			return nil, err
		}
	} else if strings.TrimSpace(i.account.Token) == "" { // Or the users token is missing
		if verbose {
			fmt.Println("Account has no authorization token")
		}
		err := i.account.RefreshSession(verbose)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest("POST", "https://aisandbox-pa.googleapis.com/v1:runImageFx", strings.NewReader(p.String()))
	if err != nil {
		if verbose {
			fmt.Println("Failed to create new request while generating image")
		}

		return nil, err
	}

	// Now lets set the authorization header in our DefaultHeader
	temp := map[string]string{"Authorization": "Bearer " + i.account.Token}
	temp = utils.MergeMap(temp, types.DefaultHeader)
	for key, value := range temp {
		req.Header.Set(key, value)
	}

	res, err := utils.Fetch(req, verbose)
	if err != nil {
		return nil, err
	}

	var parsedResponse struct {
		ImagePanels []struct {
			GeneratedImages []image.Image `json:"generatedImages"`
		} `json:"imagePanels"`
	}

	err = json.Unmarshal([]byte(res), &parsedResponse)
	if err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Failed to parse the JSON response %s: %v\n", res, err)
		}

		return nil, err
	}

	// Returning pointer to save space
	return &parsedResponse.ImagePanels[0].GeneratedImages, nil
}

func GetImageFromID(i *ImageGo, mediaID string, verbose bool) (*image.Image, error) {
	if strings.TrimSpace(mediaID) == "" {
		if verbose {
			fmt.Println("You forgot to provide mediaID?")
		}

		return nil, fmt.Errorf("please provide media id")
	}
	if i.Cookie != i.account.Cookie {
		if verbose {
			fmt.Println("Cookie don't match, lets refresh")
		}
		i.account.Cookie = i.Cookie
		err := i.account.RefreshSession(verbose)
		if err != nil {
			return nil, err
		}
	} else if strings.TrimSpace(i.account.Token) == "" { // Or the users token is missing
		if verbose {
			fmt.Println("Account has no authorization token")
		}
		err := i.account.RefreshSession(verbose)
		if err != nil {
			return nil, err
		}
	}

	body := `{"json":{"mediaKey":"` + mediaID + `"}}`
	body = url.QueryEscape(body)
	req, err := http.NewRequest("GET", "https://labs.google/fx/api/trpc/media.fetchMedia?input="+body, nil)
	if err != nil {
		if verbose {
			fmt.Println("Failed to create new request")
		}

		return nil, err
	}

	// Now lets set the authorization header in our DefaultHeader
	temp := map[string]string{"Cookie": i.account.Cookie, "Authorization": "Bearer " + i.account.Token}
	temp = utils.MergeMap(temp, types.DefaultHeader)
	for key, value := range temp {
		req.Header.Set(key, value)
	}

	res, err := utils.Fetch(req, verbose)
	if err != nil {
		return nil, err
	}

	var parsedResponse struct {
		Result struct {
			Data struct {
				JSON struct {
					Result struct {
						Image image.Image `json:"image"`
					} `json:"result"`
				} `json:"json"`
			} `json:"data"`
		} `json:"result"`
	}
	err = json.Unmarshal([]byte(res), &parsedResponse)
	if err != nil {
		if verbose {
			fmt.Println("Failed to parse response to json:", res)
		}

		return nil, err
	}

	return &parsedResponse.Result.Data.JSON.Result.Image, nil
}
