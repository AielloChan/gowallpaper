package picsource

import (
	"../tools"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"
)

const (
	NationalGeographic_API_URL = `http://www.nationalgeographic.com/photography/photo-of-the-day/_jcr_content/.gallery.json`
)

type ngObj struct {
	Items []ngItems `json:"items"`
}

type ngItems struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Sizes ngSize `json:"sizes"`
}

type ngSize struct {
	Size240  string `json:"240"`
	Size320  string `json:"320"`
	Size500  string `json:"500"`
	Size640  string `json:"640"`
	Size800  string `json:"800"`
	Size1024 string `json:"1024"`
	Size1600 string `json:"1600"`
	Size2048 string `json:"2048"`
}

// NationalGeographic provide pictures
func NationalGeographic(fileDir string, quality int) (string, error) {

	fmt.Println("[1/4]Checking directory...")
	os.MkdirAll(fileDir, 0777)

	fmt.Println("[2/4] Loadding API...")
	jsonData, err := tools.Fetch(NationalGeographic_API_URL)
	if err != nil {
		os.Exit(2)
	}

	var jsonObj ngObj
	fmt.Println("Parse json string...")
	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		fmt.Println("Parse json string failed.", err)
		return "", err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	currentItem := jsonObj.Items[r.Intn(len(jsonObj.Items))]
	var picURL string
	switch quality {
	case 0:
		picURL = currentItem.URL + currentItem.Sizes.Size1024
	case 1:
		picURL = currentItem.URL + currentItem.Sizes.Size2048
	default:
		picURL = currentItem.URL + currentItem.Sizes.Size1024
	}

	fileName := currentItem.Title + ".jpg"

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
