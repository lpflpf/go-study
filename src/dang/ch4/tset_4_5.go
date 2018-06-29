package main

import "fmt"

func uniq_str(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}

	ret := strings[:0]
	iter := 0
	for i := 0; i < len(strings); i++ {
		ret = append(ret, strings[i])
		for i < len(strings) && ret[iter] == strings[i] {
			i++
		}
		iter++
	}

	return ret
}

func main() {
	str := []string{"hello", "world", "world", "hello", "hello"}

	strings := uniq_str(str)

	fmt.Println(strings)
}
