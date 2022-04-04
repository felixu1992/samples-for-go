package main

import "fmt"

func reverse1(s *[3]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{1, 2, 3}
	reverse1(&a)
	fmt.Println(a)
}
