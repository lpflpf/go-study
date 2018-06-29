package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func join(m_slice []string) string {
	return strings.Join(m_slice, " ")
}

func string_union(m_slice []string) string {
	s := ""
	for _, val := range m_slice {
		s += val + " "
	}
	return s
}

func main() {
	var m_slice []string

	var size, _ = strconv.Atoi(os.Args[1])

	for i := 0; i < size; i++ {
		m_slice = append(m_slice, "abcdefghijklmnopqrstuvwxyz")
	}

	start := time.Now()
	join(m_slice)
	fmt.Printf("join time : %2fs elapsed\n", time.Since(start).Seconds())

	start = time.Now()
	string_union(m_slice)
	fmt.Printf("string_union : %2fs elapsed\n", time.Since(start).Seconds())

}
