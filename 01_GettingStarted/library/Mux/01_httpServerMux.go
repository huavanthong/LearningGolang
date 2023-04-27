package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Running Hello Handler")

	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)

		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(w, "Hello %s", b)
}

func main() {

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.HandleFunc("/", HelloHandler)

	// Listen for connections on all ip address (0.0.0.0)
	// port 9090
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
