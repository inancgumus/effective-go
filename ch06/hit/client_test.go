package hit

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestClientDo(t *testing.T) {
	t.Parallel()

	const wantHits, wantErrors = 10, 0
	var (
		gotHits atomic.Int64
		server  = newTestServer(t, func(_ http.ResponseWriter, _ *http.Request) {
			gotHits.Add(1)
		})
		request = newRequest(t, http.MethodGet, server.URL)
	)

	var c Client
	sum := c.Do(context.Background(), request, wantHits)
	if got := gotHits.Load(); got != wantHits {
		t.Errorf("hits=%d; want %d", got, wantHits)
	}
	if got := sum.Requests; got != wantHits {
		t.Errorf("Requests=%d; want %d", got, wantHits)
	}
	if got := sum.Errors; got != wantErrors {
		t.Errorf("Errors=%d; want %d", got, wantErrors)
	}
}

func newTestServer(tb testing.TB, h http.HandlerFunc) *httptest.Server {
	tb.Helper()
	s := httptest.NewServer(h)
	tb.Cleanup(s.Close)
	return s
}

func newRequest(tb testing.TB, method, url string) *http.Request {
	tb.Helper()
	r, err := http.NewRequest(method, url, http.NoBody)
	if err != nil {
		tb.Fatalf("newRequest(%q, %q) err=%q; want nil", method, url, err)
	}
	return r
}
