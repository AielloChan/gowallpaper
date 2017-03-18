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
	BAIDU_API_URL = `http://image.baidu.com/channel/listjson?tag1=%E5%A3%81%E7%BA%B8&tag2=%E9%A3%8E%E6%99%AF&pn=0&rn=50`
)

type baiduObj struct {
	Data []baiduData `json:"data"`
}

type baiduData struct {
	Abs   string `json:"abs"`
	Image string `json:"image_url"`
}

// Baidu provide pictures
func Baidu(fileDir string, quality int) (string, error) {

	fmt.Println("[1/4]Checking directory...")
	os.MkdirAll(fileDir, 0777)

	fmt.Println("[2/4] Loadding API...")
	jsonData, err := tools.Fetch(BAIDU_API_URL)
	if err != nil {
		os.Exit(2)
	}

	var jsonObj baiduObj
	fmt.Println("Parse json string...")
	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		fmt.Println("Parse json string failed.", err)
		return "", err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var picURL string
	switch quality {
	case 0:
		picURL = jsonObj.Data[r.Intn(len(jsonObj.Data))].Image
	case 1:
		picURL = jsonObj.Data[r.Intn(len(jsonObj.Data))].Image
	default:
		picURL = jsonObj.Data[r.Intn(len(jsonObj.Data))].Image
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
