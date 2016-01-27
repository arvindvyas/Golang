package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for key, value := range nums {
		sum += value
		fmt.Println("key: ", key, " sum is: ", sum)
	}
	fmt.Println("Final sum:", sum)

	for key := range nums {
		if key.expired() {
			delete(nums, key)
		}
	}

}
