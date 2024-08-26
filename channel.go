package main

import (
	"fmt"
)

var httpChannel = make(chan string, 3)

func getHTTPRequest(url string) {
	httpChannel <- "Going call " + url
	// fmt.Println(url)
	// httpChannel <- 1
}

func main() {
	fmt.Println("Hello Go Channel !")

	getHTTPRequest("intel.com")
	getHTTPRequest("google.com")
	getHTTPRequest("fb.com")

	fmt.Println(len(httpChannel), cap(httpChannel))
	fmt.Println(<- httpChannel)
	fmt.Println(len(httpChannel), cap(httpChannel))
	fmt.Println(<- httpChannel)
	fmt.Println(len(httpChannel), cap(httpChannel))
	getHTTPRequest("fb.com1")
	fmt.Println(<- httpChannel)
	fmt.Println(len(httpChannel), cap(httpChannel))
	fmt.Println(<- httpChannel)
	fmt.Println(len(httpChannel), cap(httpChannel))
	fmt.Println(<- httpChannel)
}