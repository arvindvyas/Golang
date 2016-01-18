package main

import "fmt"
import "os"

func main() {

	// You can get individual args with normal indexing

	fmt.Println(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
}
