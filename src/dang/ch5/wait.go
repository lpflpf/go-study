// wait
package main

import (
	"fmt"
	"time"
	"net/http
	"log"
	"errors"
)

func WaitForServer(url string err){
	const timeout = 1 * time.Minute
	dealine := time.Now().Add(timeout)
	
	errors.New("EOF")
	for tries := 0; time.Now().Before(dealine); tries ++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Fatalf("server not responding (%s); retrying...", err) 
		time.Sleep(time.Second << uint(tries))
	}
	
	return fmt.Errorf("srver %s failed to respond after %s", url, timeo)
}

func main() {
	fmt.Println("Hello World!")
}
