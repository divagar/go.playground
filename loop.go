package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 20; i++ {
		if divby3(i) == 0 && divby5(i) == 0 {
			fmt.Println("fizz buzz")
		} else if divby3(i) == 0 {
			fmt.Println("fizz")
		} else if divby5(i) == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func divby3(x int) int {
	return x % 3
}

func divby5(x int) int {
	return x % 5
}
