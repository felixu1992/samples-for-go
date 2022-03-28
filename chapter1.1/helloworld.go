package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world")
	s := os.Args
	fmt.Println(s[1:3])
}
