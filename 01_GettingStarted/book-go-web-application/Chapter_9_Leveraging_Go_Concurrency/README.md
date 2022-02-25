# Introduction
This tutorial will help you answer question below:

# 9.1 Goroutines
* [How do you use goroutines in Golang?](#Goroutines)
* [Why do you run printLetters1() case at using-goroutine, you don't see any output?(#Goroutines)
* [What purpose for waiting for goroutine?](#Waiting-for-goroutines)



## Goroutines
### Using-goroutines
To use goroutines in your code, we can use keyword "go".
```
go printNumbers1()
go printLetters1()
```
Refer to code at: [here](https://github.com/huavanthong/MasterGolang/tree/feature/chapter9/01_GettingStarted/book-go-web-application/Chapter_9_Leveraging_Go_Concurrency/9.2.1_Demonstrating_goroutines)  
If you run on the test case, we don't need to implement content of main() function because it will redirect to test code for running.  
To run test case. 
```
    go test –v
```
#### Issue-1
If you can't see output at test case 2, because the processing work being done before goroutine display the output. 
To fix this issue, you need to add a line code
```
    time.Sleep(1 * time.Millisecond)
```
If you run this code again, the last line produces a differnt result. In fact, **printNumbers2** and **printLetters2** run independently and fight to print to the screen.   
Running repeatedly will produce different results each time. If you’re using a Go version prior to Go 1.5, you might get the same results each time. Why?  
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


### Waiting-for-goroutines
Take a look at the previous part:
* Right now, you saw how goroutines are run independently.
* You also saw how the goroutines started in the program would end unceremoniously when the program ended. You got away with it by adding a time delay using the Sleep() function,
but that’s a very hacky way of handling it. Refer: [here](#Issue-1)  

Although the danger of a program ending before the goroutines can complete is less probable in any serious code (because you’ll know right away and change it), you may often encounter a need to ensure all goroutines complete before moving on to the next thing.
#### sync-package
Go provides a simple mechanism called the **WaitGroup**, which is found in the **sync** package. The mechanism is straightforward:
1. Declare a WaitGroup.
2. Set up the WaitGroup’s counter using the Add method.
3. Decrement the counter using the Done method whenever a goroutine com-
pletes its task.
4. Call the Wait method, which will block until the counter is 0.

