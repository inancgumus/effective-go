package url_test

import (
	"testing"

	"github.com/inancgumus/effective-go/ch04/url"
)

func TestParseScheme(t *testing.T) {
	const (
		rawurl     = "https://foo.com/go"
		wantScheme = "https"
		wantRest   = "foo.com/go"
		wantOk     = true
	)
	scheme, rest, ok := url.ParseScheme(rawurl)
	if scheme != wantScheme {
		t.Errorf("parseScheme(%q) scheme = %q, want %q", rawurl, scheme, wantScheme)
	}
	if rest != wantRest {
		t.Errorf("parseScheme(%q) rest = %q, want %q", rawurl, rest, wantRest)
	}
	if ok != wantOk {
		t.Errorf("parseScheme(%q) ok = %t, want %t", rawurl, ok, wantOk)
	}
}
