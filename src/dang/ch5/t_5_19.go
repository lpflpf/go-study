package main

import "fmt"

func test(t int) {
	panic(t)
}

func main() {
	fmt.Printf("%v\n", getData(100))
}

func getData(t int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = r.(int)
		}
	}()

	test(t)
	return 0
}
