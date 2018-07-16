package main

import "fmt"

const (
	mask_1 = 0x80
	mask_2 = 0xC0
	mask_3 = 0xE0
	mask_4 = 0xF0
)

func pureReverse(strings []byte) []byte {
	size := len(strings) - 1
	for i := 0; i < size; i++ {
		strings[size], strings[i] = strings[i], strings[size]
		size--
	}
	return strings
}

func reverse(strings []byte) []byte {
	return pureReverse(revereSub(strings))
}

func revereSub(strings []byte) []byte {
	for i := 0; i < len(strings); {
		if mask_4 == (strings[i] & mask_4) {
			pureReverse(strings[i : i+4])
			i += 4
			continue
		}

		if mask_3 == (strings[i] & mask_3) {
			pureReverse(strings[i : i+3])
			i += 3
			continue
		}

		if mask_2 == (strings[i] & mask_2) {
			pureReverse(strings[i : i+2])
			i += 2
			continue
		}
		i++
	}
	return strings
}

func main() {
	fmt.Printf("%s\n", string(reverse([]byte("中国你好abcHello嗯"))))
}
