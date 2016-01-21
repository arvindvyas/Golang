package main

import (
	"fmt"
	"math/big"
)

func main() {

	verybig := big.NewInt(1)

	ten := big.NewInt(10)

	for i := 0; i < 10000; i++ {
		temp := new(big.Int)
		temp.Mul(verybig, ten)
		verybig = temp
	}
	fmt.Println(verybig)
}
