package main

import "fmt"
import s "strings"

var p = fmt.Println // alias fmt.Println to shorter name
func main() {
	p("ToLower:  ", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("Index:", s.Index("test", "e"))
	p("Contains", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
}
