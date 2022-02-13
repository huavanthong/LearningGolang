package main

import (
	"html/template"
	"os"
)

type Foo struct {
	Data [9][9]int
}

func main() {
	tmpl := template.Must(template.New("example").Parse(`
    <table>
    {{range .Data}}
    <tr>
      {{range .}}<td>{{.}}</td>{{end}}
    </tr>
    {{end}}
    `))
	foo := new(Foo)

	// Assign data for array
	for i := 0; i < len(foo.Data); i++ {
		for j := 0; j < len(foo.Data); j++ {
			foo.Data[i][j] = j
		}
	}

	tmpl.Execute(os.Stdout, foo)
}
