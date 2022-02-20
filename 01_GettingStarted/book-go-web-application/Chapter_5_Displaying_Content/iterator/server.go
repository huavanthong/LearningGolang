package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/************************************************************************************************
This is a example to show how to create using contional actions in 5.3.2 book go-web-application.
This function will show you how to pass array string to template
************************************************************************************************/
// Create slice one dimensional
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	// Create slice of string to contain all days
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for day := range daysOfWeek {
		fmt.Println(day, daysOfWeek[day])
	}
	//Template
	/*
		<ul>
		{{ range . }}
		  <li>{{ . }}</li>
		{{ end}}
		</ul>
	*/
	t.Execute(w, daysOfWeek)
}

// Solution to create multiple 2d array in Golang
type Foo struct {
	Data [9][9]int
}

// Create slice two dimensional
func process2(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.New("example").Parse(`
		<table>
		{{range .Data}}
		<tr>
		{{range .}}<td>{{.}}</td>{{end}}
		</tr>
		{{end}}
    `))

	// Create a struct with array 2d in Go
	foo := new(Foo)

	// Assign value to Golang
	for i := 0; i < len(foo.Data); i++ {
		for j := 0; j < len(foo.Data); j++ {
			foo.Data[i][j] = j
		}
	}

	// Pass array to html
	t.Execute(w, foo)
	// ===> Please implement template match to array 2d at tmpl2.html
}

// Create map info and pass to template
func process3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl3.html")
	// Create slice of string to contain all days
	var countryCapitalMap map[string]string
	/* create a map*/
	countryCapitalMap = make(map[string]string)

	/* insert key-value pairs in the map*/
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	t.Execute(w, countryCapitalMap)
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
