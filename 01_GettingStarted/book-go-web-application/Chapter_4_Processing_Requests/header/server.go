package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)

}

func headersByMethods(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, "Hello World")
	r.Header.Set("Host", "domain.tld")
	h2 := r.Header
	fmt.Fprintln(w, h2)
}

func headersByMethodsDelete(w http.ResponseWriter, r *http.Request) {

	r.Header.Set("Host", "domain.tld")
	h := r.Header.Get("Host")
	fmt.Fprintln(w, h)
	r.Header.Del("Host")
	h2 := r.Header.Get("Host")
	fmt.Fprintln(w, h2, "Nothing we can get")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/headersByMethod", headersByMethods)
	http.HandleFunc("/headersByMethodDelete", headersByMethodsDelete)
	server.ListenAndServe()
}
