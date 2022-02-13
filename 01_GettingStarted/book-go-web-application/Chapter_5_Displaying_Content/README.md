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
* [How to create a function map to template engine?](#Functions)
* [What is work flow running Funcmap in Golang?](#Work-flow)
* [How many application for using function in Golang](#App)
# Context awareness
* [What is contex-aware in Golang?](#Context-awareness)
* [What is XSS attacks?](#What-is-XSS-attacks)
* [How does we depend against XSS attacks?](#Defending-against-XSS-attacks)
* [What is unescaping HTML?](#Unescaping-HTML)
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
On template, from template 1, we pass template 2 to that.
```
  {{ template "t2.html" }}
```
To pass arguement from template 1 to template 2
```
  {{ template "t2.html" . }}
```
## Arguments-Variables-Pipelines
Because this hasn't example code, we have to do some example at: [here](https://github.com/huavanthong/MasterGolang/tree/feature/chapter5/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/arg-var-pipe)
### Arguments
In template, we have a argument as arg
```
  {{ if arg }}
    some content
  {{ end }}
```
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
## Functions
To define custom functions, you need to:
1. Create a FuncMap map, which has the name of the function as the key and the actual function as the value.
2. Attach the FuncMap to the template.
For a example, we create a formatDate function and mapping to template.  
Create formatDate() function:
```
  func formatDate(t time.Time) string { ... }
```
Mapping to template
```
  funcMap := template.FuncMap{"fdate": formatDate}
  t := template.New("tmpl.html").Funcs(funcMap)
```
On template, to regconize that function, we to use 
```
  <div>The date/time is {{ . | fdate }}</div>
```
More details: [here]()
### Work-flow 
Refer to book go-web-application, chapter 5.5
### App

## Context-awareness
One of the most interesting features of the Go template engine is that the content it displays can be changed according to its context. The content that's displayed changes depending on where you place it within the documents.
This means:
- The content is HTML, it will be HTML escaped.
- If it is JavaScript, it will be JavaScript escaped.
- Go template also recognize content that's part of a URL or is a CSS style.
To create a content:
```
func process(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("tmpl.html")
  content := `I asked: <i>"What's up?"</i>`
  t.Execute(w, content)
}
```
More details: [here](https://github.com/huavanthong/MasterGolang/tree/feature/chapter5-part5.6/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/context_aware)
### What-is-XSS-attacks
A common XSS attacks is the persistent XSS vulnerability. This happens when data provided by an attacker is saved to the server and then displayed to other users as it is. 
For a example:
- a vulnerable forum site that allows its users to create posts or comments to be saved and read by other users.
- An attacker can post a comment that includes malicious JavaScript code within the <script> tag. 
- Because the forum displays the comment as is and whatever is within the <script> tag isn’t shown to the user
- the malicious code is executed with the user’s permissions but without the user’s knowledge.
The normal way to prevent this is to escape whatever is passed into the system before displaying or storing it.  
But as with most exploits and bugs, the biggest culprit is the human factor.

### Defending-against-XSS-attacks
Step 1: Now, we implement a page to submit comments from user.
- And note that, we create a action to redirect process site.
```
    <form action="/process" method="post">
      Comment: <input name="comment" type="text" size="50">
     <hr/>
     <button id="submit">Submit</button>
    </form>
```
At server.go side, we will handle to depend XSS attacks.  
Step 2: We implement a page to submit data. 
- We will check data at this /form site.
- If data is normal, we will redirect to /process site.
```
func form(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("form.html")
  t.Execute(w, nil)
}
```
Step 3: Implement a /process page to display data from redirect at /form
- At here, we will get data from request by using FormValue() method.
- If data is correct, pass to tmpl.html template 
```
func process(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("tmpl.html")
  t.Execute(w, r.FormValue("comment"))
}
```
More details: [here](https://github.com/huavanthong/MasterGolang/tree/feature/chapter5-part5.6/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/xss)
#### How to run
Open brower:
```
  http://localhost:8080/form
```
#### Test
**Case 1:** User submit a normal comment
```
  Hello World
Output:
  Hello World
  Hello World
```
**Case 2:** Submit a comment, maybe XSS attacks
```
  <script>alert('Pwnd!');</script>
Output:
  <script>alert('Pwnd!');</script>
```
As you can see, we can't send this message to /process.  
Right now, we've ready to depend XSS attacks.

### Unescaping-HTML
Suppose if you really want this behavior, meaning you want the user to enter HTML or JavaScript code that’s executable when displayed.  
Go provides a mechanism to “unescape” HTML.  
To stop the brower from protecting you from XSS attacks:
```
  w.Header().Set("X-XSS-Protection", "0")
```
### Test
**Case 1:**  If you don't stop protecting:
```
Submit a comment, maybe XSS attacks 
  <script>alert('Pwnd!');</script>
Output will display white page.
  None
```
**Case 2:**  If you stop protecting:
```
Submit a comment, maybe XSS attacks
  <script>alert('Pwnd!');</script>
Output:
  <script>alert('Pwnd!');</script>
```
Please try to do it at [here](https://github.com/huavanthong/MasterGolang/tree/feature/chapter5-part5.6/01_GettingStarted/book-go-web-application/Chapter_5_Displaying_Content/xss)
