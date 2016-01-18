package main

import "fmt"

func main() {
	var fh float32;
  fmt.Println("Enter Fahrenheit number to change it in celsius")
	fmt.Scanf("%g",&fh)
	c := (fh - 32) * 5 / 9
	fmt.Println("This is celsius form of fahrenheit you entered", c)
}
