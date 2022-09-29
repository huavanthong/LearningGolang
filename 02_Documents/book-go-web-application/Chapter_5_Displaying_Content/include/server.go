package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
}

func process1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1_pass_arg_to_t2.html", "t2.html")
	t.Execute(w, "Hello World!")
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
	http.HandleFunc("/index", index)
	http.HandleFunc("/process", process)
	http.HandleFunc("/process1", process1)
	server.ListenAndServe()
}
