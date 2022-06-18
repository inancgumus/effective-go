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
	sumUp := func() {
		total += fetch("http://foo.com")
	}
	for i := 0; i < N; i++ {
		go sumUp()
	}
	fmt.Println(total)
}
