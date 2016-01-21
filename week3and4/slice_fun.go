package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30}
	slice2 := append(slice1, 40, 50)
	fmt.Println(slice1, slice2)
}
