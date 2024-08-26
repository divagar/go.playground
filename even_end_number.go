package main

import (
	"fmt"
)

func main() {
	var a = 143612
	var b = fmt.Sprintf("%d", a)
	fmt.Printf("a = %d  b = %q typeofa = %T typeofb = %T\n", a, b, a, b)
	fmt.Println(len(b))
	fmt.Println(b[0] == b[len(b)-1])
}
