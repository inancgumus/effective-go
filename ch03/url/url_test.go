package url

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	const rawurl = "https://foo.com/go"

	u, err := Parse(rawurl)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawurl, err)
	}
	if got, want := u.Scheme, "https"; got != want {
		t.Errorf("Parse(%q).Scheme = %q; want %q", rawurl, got, want)
	}
	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawurl, got, want)
	}
	if got, want := u.Path, "go"; got != want {
		t.Errorf("Parse(%q).Path = %q; want %q", rawurl, got, want)
	}
}

func TestURLHost(t *testing.T) {
	tests := map[string]struct {
		in       string // URL.Host field
		hostname string
		port     string
	}{
		"with port":       {in: "foo.com:80", hostname: "foo.com", port: "80"},
		"with empty port": {in: "foo.com", hostname: "foo.com", port: ""},
		"without port":    {in: "foo.com:", hostname: "foo.com", port: ""},
		"ip with port":    {in: "1.2.3.4:90", hostname: "1.2.3.4", port: "90"},
		"ip without port": {in: "1.2.3.4", hostname: "1.2.3.4", port: ""},
		// Add more tests in case of a need
	}
	for name, tt := range tests {
		t.Run(fmt.Sprintf("Hostname/%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Hostname(), tt.hostname; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
		t.Run(fmt.Sprintf("Port/%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}

// ========================================================================
// #1: Repetitive tests for the same logic
/*
func TestURLPortWithPort(t *testing.T) {
	const in = "foo.com:80"

	u := &URL{Host: in}
	if got, want := u.Port(), "80"; got != want {
		t.Errorf("for host %q; got %q; want %q", in, got, want)
	}
}

func TestURLPortWithEmptyPort(t *testing.T) {
	const in = "foo.com:"

	u := &URL{Host: in}
	if got, want := u.Port(), ""; got != want {
		t.Errorf("for host %q; got %q; want %q", in, got, want)
	}
}

func TestURLPortWithoutPort(t *testing.T) {
	const in = "foo.com"

	u := &URL{Host: in}
	if got, want := u.Port(), ""; got != want {
		t.Errorf("for host %q; got %q; want %q", in, got, want)
	}
}

func TestURLPortIPWithPort(t *testing.T) {
	const in = "1.2.3.4:90"

	u := &URL{Host: in}
	if got, want := u.Port(), "90"; got != want {
		t.Errorf("for host %q; got %q; want %q", in, got, want)
	}
}

func TestURLPortIPWithoutPort(t *testing.T) {
	const in = "1.2.3.4"

	u := &URL{Host: in}
	if got, want := u.Port(), ""; got != want {
		t.Errorf("for host %q; got %q; want %q", in, got, want)
	}
}
*/

// ========================================================================
// #2: Test helpers
/*
func testPort(t *testing.T, in, wantPort string) {
	t.Helper()

	u := &URL{Host: in}
	if got := u.Port(); got != wantPort {
		t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
	}
}

func TestURLPortWithPort(t *testing.T)      { testPort(t, "foo.com:80", "80") }
func TestURLPortWithEmptyPort(t *testing.T) { testPort(t, "foo.com:", "") }
func TestURLPortWithoutPort(t *testing.T)   { testPort(t, "foo.com", "") }
func TestURLPortIPWithPort(t *testing.T)    { testPort(t, "1.2.3.4:90", "90") }
func TestURLPortIPWithoutPort(t *testing.T) { testPort(t, "1.2.3.4", "") }
*/

// ========================================================================
// #3: Table-Driven testing: Missing test names and line numbers
/*
func TestURLPort(t *testing.T) {
	tests := []struct {
		in   string // URL.Host field
		port string
	}{
		{in: "foo.com:80", port: "80"}, // with port
		{in: "foo.com:", port: ""},     // with empty port
		{in: "foo.com", port: ""},      // without port
		{in: "1.2.3.4:90", port: "90"}, // ip with port
		{in: "1.2.3.4", port: ""},      // ip without port
		// Add more tests in case of a need
	}
	for _, tt := range tests {
		u := &URL{Host: tt.in}
		if got, want := u.Port(), tt.port; got != want {
			t.Errorf("for host %q; got %q; want %q", tt.in, got, want)
		}
	}
}
*/

// ========================================================================
// #4: Ordinals as names
/*
func TestURLPort(t *testing.T) {
	tests := []struct {
		in   string // URL.Host field
		port string
	}{
		1: {in: "foo.com:80", port: "80"}, // with port
		3: {in: "foo.com:", port: ""},     // with empty port
		2: {in: "foo.com", port: ""},      // without port
		4: {in: "1.2.3.4:90", port: "90"}, // ip with port
		5: {in: "1.2.3.4", port: ""},      // ip without port
		// Add more tests in case of a need
	}
	for i := 1; i < len(tests); i++ {
		tt := tests[i]
		u := &URL{Host: tt.in}
		if got, want := u.Port(), tt.port; got != want {
			t.Errorf("test %d: for host %q; got %q; want %q", i, tt.in, got, want)
		}
	}
}
*/

// ========================================================================
// #5: Got test names
/*
func TestURLPort(t *testing.T) {
	tests := []struct {
		name string
		in   string // URL.Host field
		port string
	}{
		{
			name: "with port",
			in:   "foo.com:80", port: "80",
		},
		{
			name: "with empty port",
			in:   "foo.com:", port: "",
		},
		{
			name: "without port",
			in:   "foo.com", port: "",
		},
		{
			name: "ip with port",
			in:   "1.2.3.4:90", port: "90",
		},
		{
			name: "ip without port",
			in:   "1.2.3.4", port: "",
		},
		// Add more tests in case of a need
	}
	for _, tt := range tests {
		u := &URL{Host: tt.in}
		if got, want := u.Port(), tt.port; got != want {
			t.Errorf("%s: for host %q; got %q; want %q", tt.name, tt.in, got, want)
		}
	}
}
*/

// ========================================================================
// #6: Shuffled tests with a map
/*
func TestURLPort(t *testing.T) {
	tests := map[string]struct {
		in   string // URL.Host field
		port string
	}{
		"with port":       {in: "foo.com", port: ""},
		"with empty port": {in: "foo.com:80", port: "80"},
		"without port":    {in: "foo.com:", port: ""},
		"ip with port":    {in: "1.2.3.4:90", port: "90"},
		"ip without port": {in: "1.2.3.4", port: ""},
		// Add more tests in case of a need
	}
	for name, tt := range tests {
		u := &URL{Host: tt.in}
		if got, want := u.Port(), tt.port; got != want {
			t.Errorf("%s: for host %q; got %q; want %q", name, tt.in, got, want)
		}
	}
}
*/

// ========================================================================
// #7: Subtests without tables (with duplication problem)
/*
func TestURLPort(t *testing.T) {
	t.Run("with port", func(t *testing.T) {
		const in = "foo.com:80"

		u := &URL{Host: in}
		if got, want := u.Port(), "80"; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
	t.Run("with empty port", func(t *testing.T) {
		const in = "foo.com:"

		u := &URL{Host: in}
		if got, want := u.Port(), ""; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
	t.Run("without port", func(t *testing.T) {
		const in = "foo.com"

		u := &URL{Host: in}
		if got, want := u.Port(), ""; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
	t.Run("ip with port", func(t *testing.T) {
		const in = "1.2.3.4:90"

		u := &URL{Host: in}
		if got, want := u.Port(), "90"; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
	t.Run("ip without port", func(t *testing.T) {
		const in = "1.2.3.4"

		u := &URL{Host: in}
		if got, want := u.Port(), ""; got != want {
			t.Errorf("for host %q; got %q; want %q", in, got, want)
		}
	})
}
*/

// ========================================================================
// #8: Subtests without tables (without duplication but a tad verbose)
/*
func TestURLPort(t *testing.T) {
	testPort := func(in, wantPort string) {
		t.Helper()
		u := &URL{Host: in}
		if got := u.Port(); got != wantPort {
			t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
		}
	}
	t.Run("with port", func(t *testing.T) { testPort("foo.com:80", "80") })
	t.Run("with empty port", func(t *testing.T) { testPort("foo.com:", "") })
	t.Run("without port", func(t *testing.T) { testPort("foo.com", "") })
	t.Run("ip with port", func(t *testing.T) { testPort("1.2.3.4:90", "90") })
	t.Run("ip without port", func(t *testing.T) { testPort("1.2.3.4", "") })
}
*/

// ========================================================================
// #9: Subtests without tables (without duplication)
/*
func TestURLPort(t *testing.T) {
	testPort := func(in, wantPort string) func(*testing.T) {
		return func(t *testing.T) {
			t.Helper()
			u := &URL{Host: in}
			if got := u.Port(); got != wantPort {
				t.Errorf("for host %q; got %q; want %q", in, got, wantPort)
			}
		}
	}
	t.Run("with port", testPort("foo.com:80", "80"))
	t.Run("with empty port", testPort("foo.com:", ""))
	t.Run("without port", testPort("foo.com", ""))
	t.Run("ip with port", testPort("1.2.3.4:90", "90"))
	t.Run("ip without port", testPort("1.2.3.4", ""))
}
*/

// ========================================================================
// #10: Subtests
/*
func TestURLPort(t *testing.T) {
	var tests = map[string]struct {
		in   string // URL.Host field
		port string
	}{
		"with port":       {in: "foo.com:80", port: "80"},
		"with empty port": {in: "foo.com", port: ""},
		"without port":    {in: "foo.com:", port: ""},
		"ip with port":    {in: "1.2.3.4:90", port: "90"},
		"ip without port": {in: "1.2.3.4", port: ""},
		// Add more tests in case of a need
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("for host %q; got %q; want %q", tt.in, got, want)
			}
		})
	}
}
*/

// ========================================================================
// #11: Concise failure messages
/*
func TestURLPort(t *testing.T) {
	var tests = map[string]struct {
		in   string // URL.Host field
		port string
	}{
		"with port":       {in: "foo.com:80", port: "80"},
		"with empty port": {in: "foo.com", port: ""},
		"without port":    {in: "foo.com:", port: ""},
		"ip with port":    {in: "1.2.3.4:90", port: "90"},
		"ip without port": {in: "1.2.3.4", port: ""},
		// Add more tests in case of a need
	}
	for name, tt := range tests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}
*/

// ========================================================================
// #12: Hostname (sharing the test cases)
/*
var hostTests = map[string]struct {
	in       string // URL.Host field
	hostname string
	port     string
}{
	"with port":       {in: "foo.com:80", hostname: "foo.com", port: "80"},
	"with empty port": {in: "foo.com", hostname: "foo.com", port: ""},
	"without port":    {in: "foo.com:", hostname: "foo.com", port: ""},
	"ip with port":    {in: "1.2.3.4:90", hostname: "1.2.3.4", port: "90"},
	"ip without port": {in: "1.2.3.4", hostname: "1.2.3.4", port: ""},
	// Add more tests in case of a need
}

func TestURLHostname(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Hostname(), tt.hostname; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}

func TestURLPort(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.Port(), tt.port; got != want {
				t.Errorf("got %q; want %q", got, want)
			}
		})
	}
}
*/
