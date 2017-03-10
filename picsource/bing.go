package picsource

import (
	"../tools"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	BING_API_URL = `http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=8`
)

type bingObj struct {
	Images []bingImage `json:"images"`
}

type bingImage struct {
	URL       string `json:"url"`
	Copyright string `json:"copyright"`
}

// Bing provide pictures
func Bing(fileDir string) (string, error) {

	fmt.Println("[1/4]Checking directory...")
	os.MkdirAll(fileDir, 0777)

	fmt.Println("[2/4] Loadding API...")
	jsonData, err := tools.Fetch(BING_API_URL)
	if err != nil {
		os.Exit(2)
	}

	var jsonObj bingObj
	fmt.Println("Parse json string...")
	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		fmt.Println("Parse json string failed.", err)
		return "", err
	}

	currentItem := jsonObj.Images[0]
	picURL := "http://www.bing.com" + currentItem.URL

	fileName := currentItem.Copyright[:strings.Index(currentItem.Copyright, "，")]
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
