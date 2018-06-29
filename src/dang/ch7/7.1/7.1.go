package main

import "fmt"

import (
	"bufio"
)

type Counter struct {
	word, line int
}

func (c *Counter) Write(p []byte) (int, error) {
	k := 0
	for start := 0; start < len(p); {
		tmp, _, _ := bufio.ScanWords(p[start:], true)
		if tmp == 0 {
			break
		}
		start += tmp
		c.word += 1
		k += 1
	}

	k = 0
	for start := 0; start < len(p); {
		tmp, _, _ := bufio.ScanLines(p[start:], true)
		if tmp == 0 {
			break
		}
		start += tmp
		c.line += 1
		k += 1
	}
	return len(p), nil
}

func main() {
	var c Counter
	c.Write([]byte(`hello world	
	world1
	abcdef
	`))

	fmt.Println(c)
}
