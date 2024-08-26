package main

import (
	"fmt"
	"math"
)

func greet(who string) string {
	return "Hello " + who
}

func addsub(a int, b int) (int, int) {
	return a + b, a - b
}

func doubleAt(val []int, at int) {
	val[at] *= 2
}

func doubleIt(a int) {
	a *= 2
}

func doubleItPtr(a *int) {
	*a *= 2
}

func sqrt(a float64) (float64, error) {
	if a > 0 {
		return math.Sqrt(a), nil
	} else {
		return 0.0, fmt.Errorf("Val is less than zero - %f", a)
	}
}

func main() {
	fmt.Println(greet("Chrome"))
	fmt.Println(addsub(15, 10))

	//pass by reference
	doubleAtVal := []int{1, 2, 3, 4, 5}
	doubleAt(doubleAtVal, 2)
	fmt.Println(doubleAtVal)

	//pass by val
	doubleItVal := 15
	doubleIt(doubleItVal)
	fmt.Println(doubleItVal)

	//pass by reference/pointer
	doubleItPtr(&doubleItVal)
	fmt.Println(doubleItVal)

	//return error
	sqrtVal := 0.0
	val, err := sqrt(sqrtVal)
	fmt.Println(val, err)
}
