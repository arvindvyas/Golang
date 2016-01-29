package main

import "fmt"

type C struct {
	x float64
	int
	string
}

func main() {
	c := C{2.2, 2, "Arvind"}
	fmt.Println(c.x, c.int, c.string)
	fmt.Println("print variable c : ", c)
}
