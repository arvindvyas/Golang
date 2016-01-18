package main
import "fmt"
import "github.com/arvindvyas/Golang/Weeks1and2/ex8/ftoc"


func main() {
	var fahrenheit float32
	fmt.Println("Enter Fahrenheit Value")
        fmt.Scanf("%f", &fahrenheit)
	
        c := ftoc.Fhconvert(fahrenheit)
        fmt.Println(c,"C")
}



