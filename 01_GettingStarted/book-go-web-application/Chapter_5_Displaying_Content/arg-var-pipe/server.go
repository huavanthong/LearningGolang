package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

/************************************************************************************************
This is a example to show how to create using contional actions in 5.4 book go-web-application.
This function will show you how to pass a constant string, a variable, a function to template
************************************************************************************************/
// Using variable from library math/rand
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

// Create a const string
const hello string = "Hello World, I'm constant string"

func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl2.html")

	t.Execute(w, hello)
}

// Pass multiple varialbe to template
func process3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl3.html")

	var index int = 10
	ptr_to_index := new(int)
	ptr_to_index = &index

	// var info string = "I'm a value at index"
	t.Execute(w, ptr_to_index)
	// t.Execute(w, info)
}

/************************************************************************************************
Create main page to display all action in this project
************************************************************************************************/
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, index)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	http.HandleFunc("/process3", process3)
	server.ListenAndServe()
}
