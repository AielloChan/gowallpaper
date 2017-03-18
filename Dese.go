package main

import (
	"./picsource"
	"./win32api"
	"flag"
	"fmt"
)

const (
	FILE_DIR = "./pics"
)

func main() {
	customProvider := flag.String("provider", "unsplash", "Pictures provider. unsplash/bing/aibizhi/baidu/nationalgeographic")
	customQuality := flag.String("quality", "normal", "Pictures quality. normal/high")
	flag.Parse()

	var (
		filePath string
		err      error
		quality  int
	)

	switch *customQuality {
	case "normal":
		quality = 0
	case "high":
		quality = 1
	default:
		quality = 0
	}

	switch *customProvider {
	case "unsplash":
		filePath, err = picsource.Unsplash(FILE_DIR, quality)
	case "bing":
		filePath, err = picsource.Bing(FILE_DIR, quality)
	case "aibizhi":
		filePath, err = picsource.Aibizhi(FILE_DIR, quality)
	case "baidu":
		filePath, err = picsource.Baidu(FILE_DIR, quality)
	case "nationalgeographic":
		filePath, err = picsource.NationalGeographic(FILE_DIR, quality)
	default:
		filePath, err = picsource.Unsplash(FILE_DIR, quality)
	}
	if err != nil {
		fmt.Printf("Can't get picture from Unsplash. %s\n", err.Error())
	}

	fmt.Println("[Final] Set wallpaper...")
	if win32api.SetWallpaper(filePath) {
		fmt.Println("All done!")
	} else {
		fmt.Println("Set wallpaper failed!")
	}
}
