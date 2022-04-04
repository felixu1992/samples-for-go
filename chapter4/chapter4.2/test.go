package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(appendInt(x, 4, 5, 6))
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	z = make([]int, zlen, zlen)
	copy(z, x)
	copy(z[len(x):], y)
	return z
}
