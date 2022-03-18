// Reference: https://stackoverflow.com/questions/58736588/http-server-handlefunc-loop-on-timeout
package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var layout *template.Template
var WriteTimeout = 10 * time.Second

// var contextTimeout = 1 * time.Second

func main() {
	router := http.NewServeMux()
	server := &http.Server{
		Addr:        ":8080",
		Handler:     router,
		ReadTimeout: 5 * time.Second,
		// Step 1: set Timeout for writeResponse
		WriteTimeout: WriteTimeout + 10*time.Millisecond, //10ms Redundant time
		IdleTimeout:  15 * time.Second,
	}
	router.HandleFunc("/", home)
	server.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("responding\n")
	// Step 2: set Timeout for context as same as the time to write a Response
	ctx, _ := context.WithTimeout(context.Background(), WriteTimeout)
	// Step 3: We call WithCancel to get a copy of context, and we can cancel it with cancelFunc()
	worker, cancel := context.WithCancel(context.Background())
	// Step 4: defer to cancel if it finish the process
	defer cancel()

	var buffer string
	// Step 5: To response immediately, we use goroutine to fix it.
	go func() {
		// do something
		// Step 6: We can handle your task it here, and it wast 2s for your task.
		time.Sleep(2 * time.Second)
		buffer = "ready all response\n"

		// Step 6: We can do another thing at here, and it wast 10s for your second task.
		//do another
		time.Sleep(10 * time.Second)

		fmt.Printf("Check 1\n")
		cancel()
		fmt.Printf("worker finish\n")
	}()
	// Step 7: If cancel() is called ==> It will go into ctx.Done()
	// Questions: How do you know that
	//			1. when it will go into ctx.Done()
	//			2. When it will go into worker.Done()
	select {
	case <-ctx.Done():
		//add more friendly tips
		w.WriteHeader(http.StatusInternalServerError)
		t, _ := template.ParseFiles("layout.html")
		t.Execute(w, "world 1")
		return
	case <-worker.Done():
		w.Write([]byte(buffer))
		t, _ := template.ParseFiles("layout.html")
		t.Execute(w, "world 2")
		fmt.Printf("writed\n")
		return
	}
}
