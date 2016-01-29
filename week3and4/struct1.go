package main

import "fmt"

func Area(length, width float64) float64 {
	return length * width
}

func Perimeter(length, width float64) float64 {
	return 2 * (length + width)
}

func main() {
	length := 4.5
	width := 3.75

	area := Area(length, width)
	fmt.Println("Area is =", area)

	perimeter := Perimeter(length, width)
	fmt.Println("Permiter is =", perimeter)

}
