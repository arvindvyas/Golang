package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}

	//Here we use range to sum the
	//numbers in a slice. Arrays work
	//like this too.
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range on arrays and slices provides
	//both the index and value for each entry.
	//Above we didn't need the index, so we
	//ignored it with the blank identifier _.
	//Sometimes we actually want the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
}
