package main

import "fmt"

type counter int

func main() {
	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i, "is even number")
	}
}
