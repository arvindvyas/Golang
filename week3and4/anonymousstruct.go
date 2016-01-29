package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float64
}

func main() {
	b := B{A{1, 3}, 1.1, 2.2}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)

}
