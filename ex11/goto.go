package main

import "fmt"

func main() {
	i := 0
  Print_0_To_5:
	fmt.Println(i)

	i++

	if i > 5 {
		goto Escape
	}
	goto Print_0_To_5

  Escape:
	fmt.Println("Printed 0 to 5 loop")

}
