package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const usageText = `
Usage:
  hit [options] url
Options:`

type flags struct {
	url  string
	n, c int
}

func (f *flags) parse(s *flag.FlagSet, args []string) error {
	s.Usage = func() {
		fmt.Fprintln(s.Output(), usageText[1:])
		s.PrintDefaults()
	}

	s.Var(toNumber(&f.n), "n", "Number of requests to make")
	s.Var(toNumber(&f.c), "c", "Concurrency level")

	if err := s.Parse(args); err != nil {
		return err
	}
	if err := f.validate(s); err != nil {
		fmt.Fprintln(s.Output(), err)
		s.Usage()
		return err
	}

	f.url = s.Arg(0)

	return nil
}

// validate post-conditions after parsing the flags.
func (f *flags) validate(s *flag.FlagSet) error {
	if f.c > f.n {
		return fmt.Errorf("-c=%d: should be less than or equal to -n=%d", f.c, f.n)
	}
	if err := f.validateURL(s.Arg(0)); err != nil {
		return fmt.Errorf("url: %w", err)
	}
	return nil
}

func (f *flags) validateURL(s string) error {
	u, err := url.Parse(s)
	switch {
	case strings.TrimSpace(s) == "":
		err = errors.New("required")
	case err != nil:
		err = errors.New("parse error")
	case u.Scheme != "http":
		err = errors.New("only supported scheme is http")
	case u.Host == "":
		err = errors.New("missing host")
	}
	return err
}

// number is a natural number.
type number int

// toNumber is a convenience function for converting p to *number.
func toNumber(p *int) *number {
	return (*number)(p)
}

func (n *number) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	switch {
	case err != nil:
		err = errors.New("parse error")
	case v <= 0:
		err = errors.New("should be positive")
	}
	*n = number(v)
	return err
}

func (n *number) String() string {
	return strconv.Itoa(int(*n))
}
