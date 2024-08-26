package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	x = 11
	y = 20

	fmt.Printf("x = %v | type of x = %T \n", x, x)
	fmt.Printf("y = %v| type of y = %T \n", y, y)
	mean(x, y)
}

func mean(x int, y int) {
	fmt.Println(x, y)
}
