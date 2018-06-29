package main

import (
	"fmt"
	"os"
)

func main() {
	for key, val := range os.Args {
		fmt.Println(key, val)
	}
}
