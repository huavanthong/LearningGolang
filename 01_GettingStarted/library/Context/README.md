# Introduction
This parts will help you understand about using Context in Golang. Through that, we can deeply understand about how do you design to get a high performance in Golang.

# Table of contents
1. [Context with Value](#context-cancellation)
2. [Middleware](#middleware)
3. [Context Cancellation](#context-cancellation)
4. [Context Timeout](#context-timeout)
5. [gRPC](#grpc)
6. [OpenTelemetry](#opentelemetry)

# Questions
## About Context with Value
* [What is context with value](#context-with-value)
* [Why we use context with value in golang? Could you think what feature we can apply it?](#purpose-using-context-with-value)
* [Getting started context with value]()
## About Middleware

## About Context Cancellation

## About Context Timeout 
* [What is Timeout in Context](#context-timeout)
* [When a context timeout? Do you know what happened in a process?]

## About gRPC

## About OpenTelemetry

###########################################################################################

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

## Context Cancellation

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
## gRPC

## OpenTelemetry