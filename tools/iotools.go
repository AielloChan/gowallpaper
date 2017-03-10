package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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

func CheckPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
