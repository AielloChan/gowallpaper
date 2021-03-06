package picsource

import (
	"../tools"
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const (
	UNSPLASH_API_URL = `https://api.unsplash.com/photos/random?client_id=cb82c1ab9185e27543e298abb7c74f1a74b9f6daf461214f5669a237c5199a3c`
)

type unsplashObj struct {
	ID   string       `json:"id"`
	Urls unsplashURLs `json:"urls"`
}

type unsplashURLs struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
}

// Unsplash provide pictures
func Unsplash(fileDir string, quality int) (string, error) {

	fmt.Println("[1/4]Checking directory...")
	os.MkdirAll(fileDir, 0777)

	fmt.Println("[2/4] Loadding API...")
	jsonData, err := tools.Fetch(UNSPLASH_API_URL)
	if err != nil {
		os.Exit(2)
	}

	var jsonObj unsplashObj
	fmt.Println("Parse json string...")
	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		fmt.Println("Parse json string failed.", err)
		return "", err
	}

	var picURL string
	switch quality {
	case 0:
		picURL = jsonObj.Urls.Regular
	case 1:
		picURL = jsonObj.Urls.Full
	default:
		picURL = jsonObj.Urls.Regular
	}

	fileName := tools.GetNameFromURL(picURL)
	fileName += ".jpg"

	filePath := path.Join(tools.GetCurrentDirectory(), fileDir, fileName)
	if !tools.CheckPathExists(filePath) {
		fmt.Println("[3/4] Downlouding picture...")
		imageData, err := tools.Fetch(picURL)
		if err != nil {
			os.Exit(2)
		}

		fmt.Println("[4/4] Save file...")
		if err := tools.WriteFile(filePath, imageData); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	} else {
		fmt.Println("[3/4] File already exists...")
		fmt.Println("[4/4] Check file...")
	}

	return filePath, nil
}
