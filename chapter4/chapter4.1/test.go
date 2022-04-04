package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	//symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	////symbol1 := [...]string{"$", "€", "￡", "￥"}
	//fmt.Println(symbol)
	x := 1
	test(&x)
	fmt.Println(x)
}

func test(x *int) {
	*x = 10
}
