/*
Package hit provides facilities for making concurrent requests
to a server, aggregating and returning statistics.

# Client Example

Client's Do method can send many requests to a server.

Steps:

 1. Create a client and configure it.
 2. Create a new *[http].Request.
 3. Call the Client's Do method with the request.

Example:

	c := &hit.Client{
		C:       10,        		// makes ten parallell requests (default is the number of CPUs).
		RPS:     100,       		// throttles requests to one hundred per second.
		Timeout: 10 * time.Second,	// each request will time after ten seconds.
	}
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080", http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	// sends one million requests
	sum := c.Do(context.Background(), request, 1_000_000)
	sum.Fprint(os.Stdout)

# Do Example

Do function is a convenience around Client.Do.

The following example sends one million requests to localhost:

	sum, err := hit.Do(context.Background(), "http://localhost:8080", 1_000_000)
	if err != nil {
		log.Fatal(err)
	}
	sum.Fprint(os.Stdout)

You can customize Do as follows:

	sum, err := hit.Do(
		context.Background(), "http://localhost:8080", 1_000_000,
		hit.Timeout(time.Second),
		hit.Concurrency(10),
	)
	if err != nil {
		log.Fatal(err)
	}
*/
package hit
