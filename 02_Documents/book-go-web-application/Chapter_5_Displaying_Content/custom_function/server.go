package main

import (
	"html/template"
	"net/http"
	"time"
)

// Create a function for template
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

// API to map the internal function
func process(w http.ResponseWriter, r *http.Request) {
	// Mapping to internal function by using package template and FuncMap() method
	funcMap := template.FuncMap{"fdate": formatDate}
	// From template, we create Funcs for it, using Funcs() method
	t := template.New("tmpl.html").Funcs(funcMap)
	// Start to parse html
	t, _ = t.ParseFiles("tmpl.html")
	// Start to execute.
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
