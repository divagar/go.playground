package main

import (
	"fmt"
	"strings"
)

func main() {
	//testMap()
	countWordsInString("hello World how are you doing? hope you are doing well . Hello")
}

func countWordsInString(str string) {
	var a string
	var b []string
	//var c map[string]int
	c := make(map[string]int)

	a = strings.ToLower(str)
	b = strings.Split(a, " ")
	for i := range b {
		c[b[i]] += 1
	}

	for key, val := range c {
		fmt.Println(key, val)
	}

}

func testMap() {
	a := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
	}

	for key, val := range a {
		fmt.Println(key)
		fmt.Println(val)
		fmt.Printf("\n")
	}
}
