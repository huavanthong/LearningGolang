# Introduction
This tutorial will help you answer question below:

# 9.1 Goroutines
* [How do you use goroutines in Golang?](#Goroutines)



## Goroutines
### Using-goroutines
To use goroutines, we can use keyword "go"
```
go printNumbers1()
go printLetters1()
```
If you run on the test case, we don't need to implement content of main() function.  
To run test case  
```
    go test –v
```
If you run this code again, the last line produces a differnt result. In fact, **printNumbers2** and **printLetters2** run independently and fight to print to the screen.  Running repeatedly will produce different results each time. If you’re using a Go version prior to Go 1.5, you might get the same results each time. Why? 
This is because the default behavior in versions prior to Go 1.5 is to use just one CPU (even though you might have more than one CPU in my computer), unless stated otherwise. Since Go 1.5, the default behavior is the reverse—the Go runtime uses as many CPUs as the computer has. To use just one CPU, use this command:
```
    go test –v –cpu 1
```

### Goroutines-and-performance
Now that you know how goroutines behave, let’s consider goroutine performance.  
To run the benchmark.
```
    go test -run x -bench . –cpu 1
```