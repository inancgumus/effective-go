package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func fetch(uri string) int {
	n := rand.Intn(1000)
	fmt.Printf("fetched %d bytes\n", n)
	return n
}

func main() {
	const N = 2

	var (
		total int
		mu    sync.Mutex
	)
	sumUp := func() {
		mu.Lock()
		defer mu.Unlock()

		total += fetch("http://foo.com")
	}
	for i := 0; i < N; i++ {
		go sumUp()
	}

	mu.Lock()
	defer mu.Unlock()
	fmt.Println(total)
}
