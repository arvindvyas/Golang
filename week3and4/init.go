package main

import "fmt"

const Hello = "Hi"

var throne string

func init() {

	throne = "My Loard"

}

func main() {
	fmt.Println(Hello, throne)

}
