package main

import "fmt"

func main() {
	l1 := []string{"a", "b", "c"}
	Add(l1)
	fmt.Println(Count(l1))
	fmt.Println(m)
	fmt.Println("---")
	Add(l1)
	fmt.Println(Count(l1))
	fmt.Println(m)
	fmt.Println("---")
	fmt.Println(Count([]string{"1"}))
}

var m = make(map[string]int)

func Add(list []string) {
	m[k(list)] += 1
}

func Count(list []string) int {
	return m[k(list)]
}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
