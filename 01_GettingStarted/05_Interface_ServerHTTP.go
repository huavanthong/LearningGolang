

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

/*
For brevity, let's ignore POSTs and assume HTTP requests are always GETs;
that simplification does not affect the way the handlers are set up.
Here's a trivial implementation of a handler to count the number of times the page is visited.
*/

// Simple counter server.
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

// More detail:
//             https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/handler