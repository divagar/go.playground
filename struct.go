package main

import (
	"fmt"
)

type User struct {
	Firstname string
	Lastname  string
	Id        int
	Mobile    int
}

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

func NewSquare(x int, y int, len int) (*Square, error) {
	s := &Square{
		Center: Point{x, y},
		Length: len,
	}
	return s, nil
}
func (sq *Square) Move(x1 int, y1 int) {
	sq.Center.Move(x1, y1)
}

func (sq *Square) Area() int {
	return sq.Length * sq.Length
}

func main() {
	fmt.Println("Struct example")
	//user stuct
	u1 := User{"Divagar", "Mohandass", 11301084, 9535583423}
	fmt.Println(u1)

	u2 := User{
		Firstname: "D1",
		Lastname:  "M1",
		Id:        1130,
		Mobile:    9898}
	fmt.Println(u2)

	//square stuct
	sq, err := NewSquare(1, 1, 15)
	sq.Move(2, 2)
	fmt.Println(sq, err)

}
