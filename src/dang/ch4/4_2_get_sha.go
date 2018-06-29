package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	cryptoType := flag.String("type", "sha256", "crypto type.")
	str := flag.String("str", "", "crypto string")
	flag.Parse()

	switch *cryptoType {
	case "sha256":
		fmt.Printf("[% X]", sha256.Sum256([]byte(*str)))
	case "sha384":
		fmt.Printf("[% X]", sha512.Sum384([]byte(*str)))
	case "sha512":
		fmt.Printf("[% X]", sha512.Sum512([]byte(*str)))
	}
}
