package main

import "fmt"

func rotate(s []int) {
	max := len(s) - 1
	min := 0

	for ; max > min; min++,max-- {
		fmt.Printf("%d, %d\n", min, max)
		s[max], s[min] = s[min], s[max]
		max--
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rotate(s)

	fmt.Println(s)
}
