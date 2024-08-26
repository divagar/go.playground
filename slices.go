package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Slice")
	a := []string{"hello", "world", "how", "are", "you"}

	fmt.Printf("my slice is %v\n", a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("\n")
	for i := range a {
		fmt.Println(a[i])
	}
}
