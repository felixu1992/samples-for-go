package main

import "fmt"

func main() {
	num(2)
}

func num(x int) {
	switch {
	case x > 0:
		fmt.Println("case one")
	default:
		fmt.Println("case default")
	case x > 1:
		fmt.Println("case two")
	}
}
