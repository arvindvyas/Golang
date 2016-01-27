package main

import "fmt"

func main() {
	for pos, char := range "arvind@vyas" {
		fmt.Printf("character '%c' starts at byte position %d \n", char, pos)
	}

}
