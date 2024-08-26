package main

import (
	"fmt"
	"net/http"
)

func getContentType(url string) (string, string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return url, "", fmt.Errorf("Failed to make http get call")
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return url, "", fmt.Errorf("Failed to get http content ")
	} else {
		return url, contentType, nil
	}
}


func main() {
	fmt.Println("Go routines..")

	urls := []string{
		"https://intel.com",
		"https://google.com",
		"https://api.github.com",
	}

	// var wg sync.WaitGroup
	// for _, url := range urls {
	// 	wg.Add(1)
	// 	fmt.Println(url)
	// 	// fmt.Println(getContentType(url))
	// 	go func(url string) {
	// 		fmt.Println(getContentType(url))
	// 		wg.Done()
	// 	}(url)
	// }
	// wg.Wait()
	for _, url := range urls {
		fmt.Println(url)
		fmt.Println(getContentType(url))
	}
}