package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Can't connect to url: %s \n", url)
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("Content length: %d bytes.\n", resp.ContentLength)

	fmt.Println("Donloading...")
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
