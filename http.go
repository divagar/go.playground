package main

import (
	"fmt"
	"net/http"
)

func getContentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("Failed to make http get call")
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return "", fmt.Errorf("Failed to get http content ")
	} else {
		return contentType, nil
	}
}

func main() {
	fmt.Println("Hello web !")

	fmt.Println(getContentType("https://intel.com"))
}
