package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i < 8; i++ {
		fmt.Printf("x = %d e^x = %8.3f\n", i, math.Exp(float64(i)))
	}
	//a := `fasdfasdf`
}
