# Introduction
This parts will help you understand about using Context in Golang. Through that, we can deeply understand about how do you design to get a high performance in Golang.

# Table of contents
1. [Context with Value](#context-cancellation)
2. [Middleware](#middleware)
3. [Context Cancellation](#context-cancellation)
4. [Context Timeout](#context-timeout)

### Reference:
- Context usage: [here](https://dev.to/gopher/getting-started-with-go-context-l7g)
- Overview about work-flow in context: [here](https://www.sohamkamani.com/golang/context-cancellation-and-values/)
- Understand deeply about pipeline in Context: [here](https://go.dev/blog/pipelines)

# Questions
## About Context with Value
* [What is context with value](#context-with-value)
* [Why we use context with value in golang? Could you think what feature we can apply it?](#purpose-using-context-with-value)
* [Getting started context with value]()
## About Middleware
* [What is middleware?](#middleware)
* [Why the middleware can understand some processing running in the backgroud?](#middleware)
* [Do you have any idea for implementing it in middleware?](#idea-for-usage-of-context-in-middleware)
## About Context Cancellation
* [What is Context Cancellation?](#context-cancellation)
* [What purpose for using cancellation?](#getting-started-context-cancellation)

## About Context Timeout 
* [What is Timeout in Context](#context-timeout)
* [When a context timeout? Do you know what happened in a process? How do you select the happened case?](#getting-started-context-timeout)
* [What is diffrence between timeout in this, and timeout pattern in microservice?]()

===========================================================================================================================================

## Context with Value
Context with value is a simple word - you can pass a value to context running in the backgroud, and another function can use it.

### Purpose using context with value
There are many purpose using context with value. However, one of the most common uses for contexts is to share data. Through that, we can use:
1. When you have multiple functions and you want to share data between them
2. or use request scoped values.

### Getting Started Context with Value
Import **context package**
```
import "context"
```

To create a context
```
    ctx := context.Background()
```

To add a value to context, we can use ctx.WithValue()
```
    context.WithValue(ctx, "key", "test-value")
```

To read a value from context, we can use ctx.Value()
```
    ctx.Value("key")
```
**Output**
```
> fmt.print(ctx.Value("key"))
test-value
```
More details: [here](https://github.com/huavanthong/MasterGolang/blob/feature/context/01_GettingStarted/library/Context/context-with-value.go)

## Middleware
A great example and use case for request scoped values is working with middlewares in web request handlers. 
**Importance**:
- The type <http.Request> contains a context which can carry scoped values throughout the HTTP pipeline.
- It is very common to see code where middlewares are added to the HTTP pipeline and the results of the middlewares are added to the <http.Request> context. 
- This is a very useful technique as you can rely on something you know happened to in your pipeline already on later stages.
### Idea for usage of context in Middleware
```
Step 1: we can design a common key using entire program, and we add this key to context in backgroud.
Step 2: For other handler functions, we can get out this key, and using in common purpose.
```
### Getting Started with Middleware
Set a uuid at common function
```
    uuid := uuid.New()
    r = r.WithContext(context.WithValue(r.Context(), "uuid", uuid))
```

In another handler, we get out this key.
```
    uuid := r.Context().Value("uuid")
    log.Printf("[%v] Returning 200 - Healthy", uuid)
```
More details: [here](https://github.com/huavanthong/MasterGolang/blob/feature/context/01_GettingStarted/library/Context/middleware.go)
## Context Cancellation
Another very useful feature of context in golang is cancelling things that are related. This is very important when you want to propagate your cancellation. It’s a good practice to propagate the cancellation signal when you receive one. Let’s say you have a function where you start tens of goroutines. That main function waits for all goroutines to finish or a cancellation signal before proceeding. If you receive the cancellation signal you may want to propagate it to all your goroutines, so you don’t waste compute resources. If you share the same context among all goroutines you can easily do that.
* In a simple, we cancel() happened, we want all goroutines are running must be stopped.
### Ideas when using Context Cancellation
- We want to get a high performance. We want it run faster.
- The questions put out that how we can do it?
===> Yes, we can do it if you use a cancellation smarted
### Getting Started Context Cancellation
If you want to call a function to get a value faster, we can use WithCancel().
```
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 100 {
			break
		}
	}
```
More details: [here](https://github.com/huavanthong/MasterGolang/blob/feature/context/01_GettingStarted/library/Context/withCancel.go)
- [example2](https://github.com/huavanthong/MasterGolang/blob/feature/context/01_GettingStarted/library/Context/context-cancellation.go))
- To understand it deeply: [here](https://www.meisternote.com/app/note/p2AEeT5NTyfp/context-cancellation)
## Context Timeout
Context timeout is a simple word - you can set a timeout running in the backgroud, if time is expired, it will end this process using this context.

### Purpose using Context Timeout
In web application design, we can use timeout to avoid many abnormal case. For a example:
1. Use timeout to design Timeout pattern, in Microservice. Refer: [here](https://www.meisternote.com/app/note/0gdFcuDdHd3p/timeouts)
2. ...

### Getting Started Context Timeout
To set a timeout for context, we can use ctx.
```
const shortDuration = 1 * time.Millisecond

ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
```

When a timeout is set, we have the different case to select it
```
    // Select case happened on context with Timeout
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

Note:
    - shortDuration = 1ms < 1s:
        + Output: prints "context deadline exceeded"
    - shortDuration = 2000ms > 1:
        + Output: prints "overslept"
```
More details: [here](https://github.com/huavanthong/MasterGolang/blob/feature/context/01_GettingStarted/library/Context/withTimeout.go)
