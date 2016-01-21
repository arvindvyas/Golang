package main

import "fmt"

func f() {
	for i := 0; i < 5; i++ {
		defer func(n int) { fmt.Println("Bye", n) }(i)
		fmt.Println(i)
	}

}

func main() {

	f()

}
