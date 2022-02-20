package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

/************************************************************************************************
This is a example to show how to create using contional actions in 5.3.1 book go-web-application.
This function will show you how to pass a constant string, a variable, a function to template
************************************************************************************************/
// Using variable from library math/rand
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	// Access time at hardawre
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
