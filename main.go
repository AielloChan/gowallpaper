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

	provider := flag.String("provider", "unsplash", "Pictures provider. unsplash/")

	var filePath string

	switch *provider {
	case "unsplash":
		var err error
		filePath, err = picsource.Unsplash(FILE_DIR)
		if err != nil {
			fmt.Printf("Can't get picture from Unsplash. %s\n", err.Error())
		}
	}

	fmt.Println("[Final] Set wallpaper...")
	if win32api.SetWallpaper(filePath) {
		fmt.Println("All done!")
	} else {
		fmt.Println("Set wallpaper failed!")
	}
}
