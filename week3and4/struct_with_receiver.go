package main

import "fmt"

type Square struct {
	X    int
	name string
}

func (s Square) Area() int {
	return s.X * s.X

}

func main() {
	r1 := Square{5, "My Square"}
	fmt.Println("Square is", r1)
	fmt.Println("Square area is :", r1.Area())
}
