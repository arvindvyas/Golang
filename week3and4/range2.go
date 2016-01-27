package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, value := range nums {
		sum += value
		fmt.Println("Here is loop sum :", sum)
	}
	fmt.Println("Final sum:", sum)
}
