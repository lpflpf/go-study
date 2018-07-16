package main

import "bytes"
import "fmt"

func main() {
	var str = "12312312312"
	fmt.Println(str)
	fmt.Println(comma(str))
}

func comma(s string) string {
	var buffer bytes.Buffer
	size := len(s)
	reminder := size % 3

	if reminder > 0 {
		buffer.WriteString(s[:reminder])
	}

	for iter := reminder; iter < size; iter += 3 {
		buffer.WriteString(",")
		buffer.WriteString(s[iter : iter+3])
	}

	return buffer.String()
}
