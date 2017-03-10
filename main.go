package main

import (
	"./win32api"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	API_URL  = `https://api.unsplash.com/photos/random?client_id=cb82c1ab9185e27543e298abb7c74f1a74b9f6daf461214f5669a237c5199a3c`
	FILE_DIR = "./pics"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Can't connect to url: %s \n", url)
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("Content length: %d bytes.\n", resp.ContentLength)

	fmt.Println("Donloading..")
	fmt.Printf("%9d byte", 0)
	var payload []byte
	var backKeyStr string
	for i := 0; i < 15; i++ {
		backKeyStr += " "
	}
	for {
		buf := make([]byte, 1024)
		switch nr, err := resp.Body.Read(buf[:]); true {
		case nr < 0:
			fmt.Printf("Can't read content from url: %s \n", url)
			return nil, err
		case nr == 0: // EOF
			fmt.Print("\n")
			return payload, nil
		case nr > 0:
			payload = append(payload, buf[:nr]...)
			fmt.Print(backKeyStr)
			fmt.Printf("%9d bytes", len(payload))
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func CheckPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteFile(filePath string, data []byte) error {
	fileHandler, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Create file %s error!\n", filePath)
		return err
	}
	defer fileHandler.Close()

	fileHandler.Write(data)
	return nil
}

func GetNameFromURL(url string) string {
	slashIndex := strings.LastIndex(url, "/")
	questionIndex := strings.Index(url, "?")
	if questionIndex == -1 {
		return url[slashIndex+1:]
	} else {
		return url[slashIndex+1 : questionIndex]
	}
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {

	fmt.Println("Start load API...")
	jsonData, err := Fetch(API_URL)
	if err != nil {
		os.Exit(2)
	}

	type Urls struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
	}

	type Unsplash struct {
		ID   string `json:"id"`
		Urls Urls   `json:"urls"`
	}

	var jsonObj Unsplash

	fmt.Println("Parse json string...")
	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		fmt.Println("Parse json string failed.", err)
		os.Exit(2)
	}

	fileUrl := jsonObj.Urls.Full

	fmt.Println("Creating dir...")
	os.MkdirAll(FILE_DIR, 0777)

	fmt.Println("Start downloud picture...")
	imageData, err := Fetch(fileUrl)
	if err != nil {
		os.Exit(2)
	}
	fileName := GetNameFromURL(fileUrl)
	fileName += ".jpg"
	filePath := path.Join(GetCurrentDirectory(), FILE_DIR, fileName)

	fmt.Println("Save file...")
	if err := WriteFile(filePath, imageData); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println("Set wallpaper...")
	if win32api.SetWallpaper(filePath) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Failed!")
	}
}
