package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[0:5]
	fmt.Println(slice)
	fmt.Println("Length is:", len(slice))
	fmt.Println("Capacity is:", cap(slice))

	slice = slice[2:4]
	fmt.Println(slice)
	fmt.Println("Length is:", len(slice))
	fmt.Println("Capacity is:", cap(slice))

}
