package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total

}

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(add(2, 2, 2))

}
