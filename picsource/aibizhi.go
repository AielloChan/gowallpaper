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
	AIBIZHI_API_URL = `http://api.lovebizhi.com/macos_v4.php?a=category&tid=2&uuid=436e4ddc389027ba3aef863a27f6e6f9&retina=1&bizhi_width=1920&bizhi_height=1200&client_id=1008`
)

type aibizhiObj struct {
	Data []aibizhiData `json:"data"`
}

type aibizhiData struct {
	Image aibizhiImage `json:"image"`
}

type aibizhiImage struct {
	Diy string `json:"diy"`
}

// Aibizhi provide pictures
func Aibizhi(fileDir string, quality int) (string, error) {

	fmt.Println("[1/4]Checking directory...")
	os.MkdirAll(fileDir, 0777)

	fmt.Println("[2/4] Loadding API...")
	jsonData, err := tools.Fetch(AIBIZHI_API_URL)
	if err != nil {
		os.Exit(2)
	}

	var jsonObj aibizhiObj
	fmt.Println("Parse json string...")
	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		fmt.Println("Parse json string failed.", err)
		return "", err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var picURL string
	switch quality {
	case 0:
		picURL = jsonObj.Data[r.Intn(len(jsonObj.Data))].Image.Diy
	case 1:
		picURL = jsonObj.Data[r.Intn(len(jsonObj.Data))].Image.Diy
	default:
		picURL = jsonObj.Data[r.Intn(len(jsonObj.Data))].Image.Diy
	}

	fileName := tools.GetNameFromURL(picURL)

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
