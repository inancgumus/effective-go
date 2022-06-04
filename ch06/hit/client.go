package hit

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

// Client makes HTTP requests and returns an aggregated performance
// result. The fields should not be changed after initializing.
type Client struct {
	C       int           // C is the concurrency level (default: # of CPUs)
	RPS     int           // RPS throttles the requests per second
	Timeout time.Duration // Timeout per request
}

// Option allows changing Client's behavior.
type Option func(*Client)

// Concurrency changes Client's concurrency level.
func Concurrency(n int) Option {
	return func(c *Client) { c.C = n }
}

// Timeout changes Client's timeout per request.
func Timeout(d time.Duration) Option {
	return func(c *Client) { c.Timeout = d }
}

// Do sends n GET requests to the url using as many goroutines as the
// number of CPUs on the machine and returns an aggregated result.
func Do(ctx context.Context, url string, n int, opts ...Option) (*Result, error) {
	r, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("new http request: %w", err)
	}
	var c Client
	for _, o := range opts {
		o(&c)
	}
	return c.Do(ctx, r, n), nil
}

// Do sends n HTTP requests and returns an aggregated result.
func (c *Client) Do(ctx context.Context, r *http.Request, n int) *Result {
	if c == nil {
		panic("hit.Do: Client is nil")
	}
	t := time.Now()
	sum := c.do(ctx, r, n)
	return sum.Finalize(time.Since(t))
}

func (c *Client) do(ctx context.Context, r *http.Request, n int) *Result {
	p := produce(ctx, n, func() *http.Request {
		return r.Clone(ctx)
	})
	if c.RPS > 0 {
		p = throttle(ctx, p, time.Second/time.Duration(c.RPS*c.concurrency()))
	}
	var (
		sum    Result
		client = c.client()
	)
	defer client.CloseIdleConnections()
	for result := range split(p, c.concurrency(), c.send(client)) {
		sum.Merge(result)
	}
	return &sum
}

func (c *Client) send(client *http.Client) SendFunc {
	return func(r *http.Request) *Result {
		return Send(client, r)
	}
}

func (c *Client) client() *http.Client {
	return &http.Client{
		Timeout: c.Timeout,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: c.concurrency(),
		},
	}
}

func (c *Client) concurrency() int {
	if c.C > 0 {
		return c.C
	}
	return runtime.NumCPU()
}
