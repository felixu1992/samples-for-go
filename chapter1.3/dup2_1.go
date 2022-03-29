package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var duplicate []string
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		filename := duplicateFilename(f)
		if filename != "" {
			duplicate = append(duplicate, filename)
		}
		f.Close()
	}
	for _, n := range duplicate {
		fmt.Printf("%s\n", n)
	}
}

func duplicateFilename(f *os.File) string {
	input := bufio.NewScanner(f)
	counts := make(map[string]int)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if counts[text] > 1 {
			return f.Name()
		}
	}
	return ""
}
