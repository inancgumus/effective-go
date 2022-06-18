package main

import (
	"fmt"
	"math/rand"
)

func fetch(uri string) int {
	n := rand.Intn(1000)
	fmt.Printf("fetched %d bytes\n", n)
	return n
}

func main() {
	const N = 2

	var total int
	sumUp := func(c chan<- int, n int) {
		for ; n > 0; n-- {
			c <- fetch("http://foo.com")
		}
	}

	c := make(chan int)

	for i := 0; i < N; i++ {
		go sumUp(c, rand.Intn(5))
	}
	for n := range c {
		total += n
	}
	fmt.Println(total)
}
