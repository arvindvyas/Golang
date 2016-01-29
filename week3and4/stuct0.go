package main

import "fmt"

type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	r := Rectangle{2.3, 3.2}
	area := r.Area()
	fmt.Println("The area of Rectangle is :", area)

	perimeter := r.Perimeter()
	fmt.Println("Perimeter is = ", perimeter)
}
