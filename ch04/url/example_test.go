package url_test

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/inancgumus/effective-go/ch04/url"
)

func ExampleURL() {
	u, err := url.Parse("http://foo.com/go")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	fmt.Println(u)
	// Output:
	// https://foo.com/go
}

func ExampleURL_fields() {
	u, err := url.Parse("https://foo.com/go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u)
	// Output:
	// https
	// foo.com
	// go
	// https://foo.com/go
}

func ExamplePerm() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, v := range r.Perm(3) {
		fmt.Println(v)
	}
	// Unordered output:
	// 2
	// 0
	// 1
}
