package main

import (
	"fmt"
)

func main() {
	str := "hello chromebook! it is a cloud based notebook"
	fmt.Printf("str = %v | str[0] = %v | typeof str = %T | typeof str[0] = %T\n", str, str[0], str, str[0])
}
