package main

import "fmt"

// this struct is visible only in this package
// as it starts with small letter

type myStructName struct {
	// variable starts with small letter so not visible outside the package
	x int
	y string
}

func main() {
	msn := myStructName{2, "Vyas"}
	msn.x = 4
	fmt.Println(msn.x)
}
