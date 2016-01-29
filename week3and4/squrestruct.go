package main

import "fmt"

// this struct is visible outside the package as it start with capital letter

type Square struct {
	// variable start with capital letter, can visible outside the package
	X    int
	name string
}

func main() {
	sqr1 := Square{10, "Go Outside"}
	fmt.Println("Square sqr1 is:", sqr1)

	// initialize value by variable name in any order
	sqr2 := Square{name: "Arvind", X: 89}
	fmt.Println("Square sqr2 is:", sqr2)

	// get pointer to an instance with new keyword
	ps := new(Square)

	// set value using notation by derefrencing pointer
	(*ps).X = 30

	// set value using same as above
	ps.name = "Arvind Vyas"

	fmt.Println("Square ps is   ", *ps)

	ps1 := &sqr1
	fmt.Println("assing addressof sqr1 to ps1", ps1)
	ps1.X = 22

	fmt.Println("This new value of x  :", ps1.X)
	fmt.Println("Square sqr1 is :", sqr1)

	ps2 := sqr1
	ps2.X = 11

	fmt.Println("This is new value of x :", ps2.X)
}
