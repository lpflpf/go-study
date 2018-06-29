package main

import "fmt"
import "unicode"

func del_empty(strings []rune) []rune {
	ret := strings[:0]
	size := len(strings)
	iter := 0
	tag := false

	for i := 0; i < size; i++ {
		if !unicode.IsSpace(strings[i]) {
			ret = append(ret, strings[i])
			iter++
			tag = false
		} else if tag {
			continue
		} else {
			ret = append(ret, ' ')
			iter++
			tag = true
		}
	}
	return ret
}

func main() {
	// unicode.IsSpace

	strings := []rune("中国    ，china")
	fmt.Printf("%v", string(del_empty(strings)))
}
