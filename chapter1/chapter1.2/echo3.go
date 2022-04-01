package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], "/"))
	//fmt.Println(os.Args[1:])
	//fmt.Println(os.Args[0])
	for i, arg := range os.Args {
		index := i + 1
		fmt.Printf("第%d个参数为: "+arg, index)
		fmt.Println()
	}
}
