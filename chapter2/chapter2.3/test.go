package main

import (
	"fmt"
)

func main() {
	//var s string
	//fmt.Printf(s)
	x := 1
	p := &x
	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
	fmt.Println(*&x)
}
