package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/inancgumus/effective-go/ch06/hit"
)

const bannerText = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/
`

func main() {
	if err := run(flag.CommandLine, os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error occurred:", err)
		os.Exit(1)
	}
}

func banner() string { return bannerText[1:] }

func run(s *flag.FlagSet, args []string, out io.Writer) error {
	f := &flags{
		n: 100,
		c: runtime.NumCPU(),
	}
	if err := f.parse(s, args); err != nil {
		return err
	}
	fmt.Fprintln(out, banner())
	fmt.Fprintf(out, "Making %d requests to %s with a concurrency level of %d.\n",
		f.n, f.url, f.c)
	if f.rps > 0 {
		fmt.Fprintf(out, "(RPS: %d)\n", f.rps)
	}

	const timeout = time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	defer stop()

	request, err := http.NewRequest(http.MethodGet, f.url, http.NoBody)
	if err != nil {
		return err
	}
	c := &hit.Client{
		C:       f.c,
		RPS:     f.rps,
		Timeout: 10 * time.Second,
	}
	sum := c.Do(ctx, request, f.n)
	sum.Fprint(out)

	if err := ctx.Err(); errors.Is(err, context.DeadlineExceeded) {
		return fmt.Errorf("timed out in %s", timeout)
	}
	return nil
}
