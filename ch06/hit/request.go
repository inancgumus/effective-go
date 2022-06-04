package hit

import (
	"io"
	"net/http"
	"time"
)

// SendFunc is the type of the function called by Client.Do
// to send an HTTP request and return a performance result.
type SendFunc func(*http.Request) *Result

// Send an HTTP request and return a performance result.
func Send(c *http.Client, r *http.Request) *Result {
	t := time.Now()

	var (
		code  int
		bytes int64
	)
	response, err := c.Do(r)
	if err == nil {
		code = response.StatusCode
		bytes, err = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()
	}

	return &Result{
		Duration: time.Since(t),
		Bytes:    bytes,
		Status:   code,
		Error:    err,
	}
}
