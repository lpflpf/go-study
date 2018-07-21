package main

import (
	"fmt"
	"sync"
)

var n sync.WaitGroup

func test1() {
	defer n.Done()
	t := make(chan int, 3)

	t <- 1
	t <- 2
	t <- 3
	close(t)

	fmt.Println(<-t)
	fmt.Println(<-t)
	fmt.Println(<-t)
	fmt.Println(<-t)
	fmt.Println(<-t)
}

func main() {
	n.Add(1)
	go test1()
	n.Wait()
}
