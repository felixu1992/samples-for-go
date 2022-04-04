package main

import "fmt"

func main() {
	m := map[string]int{
		"felixu": 18,
		"abc":    19,
	}

	var names []string
	for a, b := range m {
		names = append(names, a)
		fmt.Println(b)
	}
	fmt.Println(names)
	for i := range names {
		fmt.Println(i)
	}

	if v, ok := m["wocao"]; !ok {
		fmt.Printf("WTF: %d", v)
	}
}
