package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	var filename []string

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts := make(map[string]int)
			countLines(f, counts)
			f.Close()
			for _, c := range counts {
				if c > 1 {
					filename = append(filename, arg)
					break
				}
			}
		}
	}
	for _, filename := range filename {
		fmt.Printf("%s\n", filename)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
