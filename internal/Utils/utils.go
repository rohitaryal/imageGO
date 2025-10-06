// Package utils contains utility function to satisfy DRY
package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// FolderExists checks if a folder exists or not
// It will return false even if theres folder but
// no permissions, etc.
func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}

		fmt.Fprintf(os.Stderr, "Failed to check path %s: %v\n", path, err)
	}

	return info.IsDir()
}

func MergeMap(map1, map2 map[string]string) map[string]string {
	newMap := make(map[string]string)

	for key, value := range map1 {
		newMap[key] = value
	}

	for key, value := range map2 {
		newMap[key] = value
	}

	return newMap
}

// Fetch function performs fetch operation on provided url
func Fetch(request *http.Request, verbose bool) (string, error) {
	if verbose {
		fmt.Fprintf(os.Stdout, "Making a %s request to: %s\n", request.Method, request.URL)
	}

	client := &http.Client{Timeout: 100 * time.Second}

	response, err := client.Do(request)
	if err != nil {
		if verbose {
			fmt.Fprintf(os.Stderr, "Failed to perform %s request on %s: %v\n", request.Method, request.URL, err)
		}

		return "", err
	}

	defer response.Body.Close()

	responseText, err := io.ReadAll(response.Body)
	if err != nil {
		if verbose {
			fmt.Fprint(os.Stderr, "Failed to convert response to byte slice.\n")
		}

		return "", err
	}

	responseString := string(responseText)

	if response.StatusCode != 200 {
		if verbose {
			fmt.Fprintf(os.Stderr, "Server responded with status code %d: %s\n", response.StatusCode, responseString)
		}

		return "", fmt.Errorf("server responded with status code %d: %s", response.StatusCode, responseString)
	}

	if verbose {
		fmt.Fprintf(os.Stdout, "Server responded with status code %d: %s\n", response.StatusCode, responseString)
	}

	return responseString, nil
}
