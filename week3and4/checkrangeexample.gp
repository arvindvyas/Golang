pckage main

import "fmt"

func main() {
        for pos, char := range "a©b" {
                fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
        }
}
