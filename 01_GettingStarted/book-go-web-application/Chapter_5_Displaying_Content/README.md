To get details contents, please refer to chapter 5 at book go-web-application.  
This tutorial will help you answer question below:
# Templates and template engine
* [The difference between templates and template engine?](#Template-and-Template-Engine)
* [How to parse templates on Go?](#Parsing-templates)
* [How we can transfer variable to html?](#Parsing-templates)
* [Suppose template have an error, how we know and handle it?](#Parsing-templates)
* [What is panic error in Golang?](#Parsing-templates)
* [Could you use another way to parse template on Go? What is executing templates?](#Executing-templates)
# Actions 
* [What is actions in Golang? How we use it?](#Actions)
* [How we can create logic "if-else" in Golang?](#Contional-actions)
* [How we create for loop in Go?](#Iterator-actions)
* [How we create for loop with two dimensional array, slice or loop in maps object?](#Iterator-actions)
* [You know about actions, thinking it, could you set actions for it or not??](#Set-actions)
* [Suppose we have many file html, how can we include all tempalate to one action?](#Include-actions)
# Arguments, variables and pipelines
* [What is arguement in a template Golang?](#Arguments)
* [How we create variable and pass to template, what expereicen for you?](#Variables)
* [What is pipeline in Golang, how we use it in templates, what is benefit to use it](#Pipelines)
# Functions

# Context awareness

# Nesting templates

# Using the block action to define default templates



# Template-and-Template-Engine

## Parsing-templates
To get a template in your working directory, please prepare:
- tmpl.html
- tmpl2.html
- server.go
To parse template files, and create a parsed template struct that you can execute it later.
```
t, _ := template.ParseFiles("tmpl.html")
```

To parse multiple template files.
```
t, _ := template.ParseFiles("tmpl.html", "tmpl2.html")
```

Golang provides another mechanism to handle errors returned by parsing templates
- The Must function wraps around a function that returns a pointer to a template and an error, and panics if the error is not a nil.
- In Go, panicking refers to a situation where the normal flow of execution is stopped, and if it’s within a function, the function returns to its caller.
- The process continues up the stack until it reaches the main program, which then crashes.
```
t := template.Must(template.ParseFiles("tmpl.html"))
```

## Executing-templates
After we have a parse template struct, we can execute it
```
t.Execute(w) 
```

To transfer a variable to Execute() method
```
t.Execute(w, "Hello World") 
```

We can use ExecuteTemplate() to combine parse method.
```
t.ExecuteTemplate(w, "tmpl.html", "Hello World!")
```
## Actions
As mentioned earlier, actions are embedded commands in Go templates, placed between a set of double braces, {{ and }}. Go has an extensive set of actions, which are
quite powerful and flexible. In this section, we’ll discuss some of the important ones:  
- Conditional actions  
- Iterator actions  
- Set actions
- Include actions

### Contional-actions
This is a way to implement "if-else" logic code in Golang with syntax:
```
  {{ if arg }}
  some content
  {{ end }}
The other variant is
  {{ if arg }}
  some content
  {{ else }}
  other content
  {{ end }}
```
Get a example
```
serve.go
  func process(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("tmpl.html")
    rand.Seed(time.Now().Unix())
    t.Execute(w, rand.Intn(10) > 5)
  }
tmpl.html
  <body>
    {{ if . }}
    Number is greater than 5!
    {{ else }}
    Number is 5 or less!
    {{ end }}
  </body>
```
More details: [here](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/random_number)
### Iterator-actions
We will use iterator to display all days of week:  
```
serve.go

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}
```
At template, we use syntax: {{ range . }} --- {{ end}} to loop
```
    <ul>
    {{ range . }}
      <li>{{ . }}</li>
    {{ end}}
    </ul>
```
More details: [here](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/iterator)
### Set-actions
We can also set actions in html following:
```
server.go:
> 	t.Execute(w, "hello") ========================> Write "hello" to tempalte
tmpl.html:
  <body>
    <div>The dot is {{ . }}</div> <===================== "hello" is displayed at here.
    <div>
      {{ with "world"}} =======================> Write a string "world" to dot.
        Now the dot is set to {{ . }} <============== "World" is writed to dot.
      {{ end }}
    </div>
    <div>The dot is {{ . }} again</div> <===================== "hello" is displayed at here again after use syntax {{with "world"}}
  </body>
```
Then output will display
```
> http://localhost:8080/process
The dot is hello 
Now the dot is set to world
The dot is hello again
```
More details: [here](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/set_dot)
### Include-actions
To parse multiple template.
```
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World!")
}
```
## Arguments-Variables-Pipelines
Because this hasn't example code, we have to do some example at: [here](https://github.com/huavanthong/MasterGolang/tree/feature/chapter5/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/arg-var-pipe)
### Arguments
In template, we have a argument as arg
> {{ if arg }}
> some content
> {{ end }}
### Variables
If you use maps object, we can easy to create variable on template as sample below
```
{{ range $key, $value := . }}
   The key is {{ $key }} and the value is {{ $value }}
{{ end }}
```
**More details:** [iterator](https://github.com/huavanthong/MasterGolang/tree/feature/chapter5/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/iterator)


If you want pass multiple multiple variable, we also need to use maps as sample below:
```
serve.go
	varmap := map[string]interface{}{
		"index": 100,
		"info":  "This is a info at index 100",
	}
template.html
    <div> Index get from server:{{ .index }}</div>
    <div> Value get from server:{{ .info }}</div>
```
**More details:** [arg-var-pipe](https://github.com/huavanthong/MasterGolang/tree/feature/chapter5/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/arg-var-pipe)
### Pipelines
To create pipeline on template
```
  {{ p1 | p2 | p3 }}
```