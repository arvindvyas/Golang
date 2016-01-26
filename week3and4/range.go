package main

import "fmt"

func main() {
	nums := []int{1, 2, 4}
	sum := 0

	for key, value := range nums {
		sum += value
		fmt.Println("Key: ", key, "sum is :", sum)
	}
	fmt.Println("Final sum:", sum)
}
