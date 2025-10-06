// Package image
package image

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"time"

	Utils "github.com/rohitaryal/imageGO/internal/Utils"
	Types "github.com/rohitaryal/imageGO/pkg/Types"
)

type Image struct {
	Seed          int               `json:"seed"`
	Model         Types.Model       `json:"modelNameType"`
	Prompt        string            `json:"prompt"`
	AspectRatio   Types.AspectRatio `json:"aspectRatio"`
	MediaID       string            `json:"mediaGenerationId"`
	EncodedImage  string            `json:"encodedImage"`
	WorkflowID    string            `json:"workflowId"`
	FingerprintID string            `json:"fingerprintLogRecordId"`
}

// Save - Saves image to destination directory
// Returns saved image's absolute path or error
func (i *Image) Save(path string, verbose bool) (string, error) {
	if path == "" {
		path = "."
	}

	if !Utils.FolderExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			if verbose {
				fmt.Fprintf(os.Stderr, "Failed to create directory %s: %v\n", path, err)
			}
			return "", err
		}
	}

	decodedImage, err := base64.StdEncoding.DecodeString(i.EncodedImage)
	if err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Failed to decode the base64 encoded image: %v\n", err)
		}
		return "", err
	}

	fileName := fmt.Sprintf("%d.%s", time.Now().UnixMilli(), "png")
	absoluteFilePath := filepath.Join(path, fileName)

	err = os.WriteFile(absoluteFilePath, decodedImage, 0o664)
	if err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Failed to write decoded image data to file %s: %v\n", absoluteFilePath, err)
		}
		return "", err
	}

	return absoluteFilePath, nil
}
