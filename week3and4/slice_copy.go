package main

import "fmt"

func main() {
	slice1 := []int{10, 23, 13}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

}
