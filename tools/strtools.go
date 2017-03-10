package tools

import "strings"

func GetNameFromURL(url string) string {
	slashIndex := strings.LastIndex(url, "/")
	questionIndex := strings.Index(url, "?")
	if questionIndex == -1 {
		return url[slashIndex+1:]
	} else {
		return url[slashIndex+1 : questionIndex]
	}
}
