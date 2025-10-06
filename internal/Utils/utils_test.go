package utils_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"testing"

	utils "github.com/rohitaryal/imageGO/internal/Utils"
)

func TestFolderExists(t *testing.T) {
	folder := "temp_folder"

	// Lets make a folder first
	err := os.Mkdir(folder, 0o777)
	// Skip the test in case we failed to make a folder
	if err != nil {
		t.Skip("Failed to create folder. Skipping this test.")
	}

	status := utils.FolderExists(folder)
	if status != true {
		os.RemoveAll(folder)
		t.FailNow()
	}

	// Now lets delete the folder
	err = os.RemoveAll(folder)
	if err != nil {
		t.Skip("Failed to delete folder. Skipping this test.")
	}

	// Now lets re-check again
	status = utils.FolderExists(folder)
	if status != false {
		t.FailNow()
	}
}

func TestMergeMap(t *testing.T) {
	map1 := map[string]string{
		"name":   "rohit",
		"gender": "******************",
		"age":    "its over six thousand",
	}

	map2 := map[string]string{
		"tinder_password":  "#include<tinder.h>",
		"twitter_password": "elon_musk_sucks",
	}

	refrenceMap := map[string]string{
		"name":             "rohit",
		"gender":           "******************",
		"age":              "its over six thousand",
		"tinder_password":  "#include<tinder.h>",
		"twitter_password": "elon_musk_sucks",
	}

	mergedMap := utils.MergeMap(map1, map2)

	// Built in method `reflect.DeepEqual` to check 2 maps
	// are equal or not
	if !reflect.DeepEqual(refrenceMap, mergedMap) {
		t.FailNow()
	}
}

func TestFetchGet(t *testing.T) {
	url := "https://httpbin.org/get?name=go"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	res, err := utils.Fetch(req, false)
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	var parsedResponse struct {
		URL string `json:"url"`
	}
	err = json.Unmarshal([]byte(res), &parsedResponse)
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	if parsedResponse.URL != url {
		t.Fatal("GET FAILED: ", err)
	}
}

func TestFetchPost(t *testing.T) {
	url := "https://httpbin.org/post"
	body := "name=go"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	res, err := utils.Fetch(req, false)
	if err != nil {
		t.Fatal("POST FAILED: ", err)
	}
	var ParsedResponse struct {
		Data string `json:"data"`
	}

	err = json.Unmarshal([]byte(res), &ParsedResponse)
	if err != nil {
		t.Fatal("GET FAILED: ", err)
	}

	if ParsedResponse.Data != body {
		t.Fatal("POST FAILED: Sent and recieved body are not same.")
	}
}
