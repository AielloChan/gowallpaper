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
	provider := flag.String("provider", "unsplash", "Pictures provider. unsplash/bing")
	flag.Parse()

	var filePath string
	var err error

	switch *provider {
	case "unsplash":
		filePath, err = picsource.Unsplash(FILE_DIR)
	case "bing":
		filePath, err = picsource.Bing(FILE_DIR)
	case "aibizhi":
		filePath, err = picsource.Aibizhi(FILE_DIR)
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
