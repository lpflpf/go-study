package main

import "fmt"
import "time"
import "os"

func main() {
	abort := make(chan int, 1)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- 0
	}()
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for count := 10; count > 0; count-- {
		fmt.Println(count)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}
}
