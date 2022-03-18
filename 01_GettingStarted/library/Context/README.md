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

## About gRPC

## About OpenTelemetry

###########################################################################################

## Context with Value
Context with value is a simple word. You can pass a value to context, and another function can use it.

### Purpose using context with value
There are many purpose using context with value. However, one of the most common uses for contexts is to share data. Through that, we can use:
1. When you have multiple functions and you want to share data between them
2. or use request scoped values.

### Getting Started Context with Value
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

## Middleware

## Context Cancellation

## Context Timeout


## gRPC

## OpenTelemetry