To get details contents, please refer to chapter 5 at book go-web-application.  
This tutorial will help you answer question below:
# Templates and template engine
* [The difference between templates and template engine?](#Template-and-Template-Engine)
* [How to parse templates on Go?](#Parsing-templates)
* [How we can transfer variable to html?](#Parsing-templates)
* [Could you use another way to parse template on Go? What is executing templates?](#Executing-templates)
# Actions 
* [What is actions in Golang? How we use it?](#Actions)
* [How we can create logic "if-else" in Golang?](#Contional-actions)
* [How we create for loop in Go?](#Iterator-actions)
* [You know about actions, thinking it, could you set actions for it or not??](#Set-actions)



## Template-and-Template-Engine

## Parsing-templates
To get a template in your working directory, please prepare:
- tmpl.html
- server.go
```
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "hello") 
}
```
How we transfer string "hello" to our template. We use dot "." sign to implement on template as:
```
```


## Executing-templates
We can use ExecuteTemplate() method to combine the above actions:
```

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
More details: [here](
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
