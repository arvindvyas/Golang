package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"Arvind", 27})
	fmt.Println(Person{name: "Neetu", age: 27})

	s := Person{"Subh", 1}
	fmt.Println(s)

	r := &Person{"Anshul", 3} //Pointer
	fmt.Println(r)

	r.age = 4       // change the value
	fmt.Println(*r) // Pointer de referencing

	// Another example of Pointers

	p := &Person{}
	p.name = "Hulk"
	p.age = 10
	fmt.Println(*p)
	fmt.Println("with out refrence", p)
}
