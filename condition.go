package main

import (
	"fmt"
)

func ifcondition(x int) {
	if x == 10 {
		fmt.Println("X is equal to 10")
	} else if x > 10 {
		fmt.Println("X is greater than 10")
	} else {
		fmt.Println("X is less than 10")
	}
}

func switchcondition(x int) {
	switch x {
	case 10:
		fmt.Println("X is 10")
	case 15:
		fmt.Println("X is 15")
	case 20:
		fmt.Println("X is 20")
	default:
		fmt.Printf("Default: X is %v", x)
	}
}

func main() {
	ifcondition(10)
	switchcondition(20)
}
