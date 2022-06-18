package main

import (
	"context"
	"fmt"
	"time"
)

func logURI(ctx context.Context, uri string) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("fetch: cannot log")
	case <-time.After(3 * time.Second):
		fmt.Printf("fetch: %s\n", uri)
		// somehow cancel the logging operation
	}
}

func fetch(ctx context.Context, uri string) int {
	go logURI(ctx, uri)

	select {
	case <-ctx.Done():
		return 0
	case <-time.After(5 * time.Second):
		return 1000
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := make(chan int, 1)
	go func() {
		c <- fetch(ctx, "http://foo.com")
	}()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case n := <-c:
		fmt.Println(n)
	}
	time.Sleep(time.Second)
}
