package main

import (
	"fmt"
	"time"
)

func fetch(uri string) int {
	<-time.After(5 * time.Second)
	return 1000
}

func main() {
	c := make(chan int)
	go func() {
		c <- fetch("http://foo.com")
	}()
	select {
	case <-time.After(time.Second):
		fmt.Println("timeout")
	case n := <-c:
		fmt.Println(n)
	}
	time.Sleep(6 * time.Second)
	panic("show me the goroutines")
}
