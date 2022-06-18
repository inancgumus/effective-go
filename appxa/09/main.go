package main

import (
	"context"
	"fmt"
	"time"
)

func fetch(ctx context.Context, uri string) int {
	select {
	case <-ctx.Done():
		return 0
	case <-time.After(5 * time.Second):
		return 1000
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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
	time.Sleep(6 * time.Second)
	panic("show me the goroutines")
}
