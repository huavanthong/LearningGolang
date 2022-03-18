# Introduction
This tutorial will help you answer question below:

# Goroutines
* [How do you use goroutines in Golang?](#Goroutines)
* [Why do you run printLetters1() case at using-goroutine, you don't see any output?](#Issue-1)
* [Is that fair if TestGoPrint1 is set delay time 1s while TestPrint1 run quickly?](#Issue-2)
# Waiting for Goroutines
* [What purpose for waiting for goroutine?](#Waiting-for-goroutines)
* [What is work-flow for using waitGroup?](#Work-flow-using-WaitGroup)
* [What happens if you forget to decrement the counter in one of the goroutines?](#Issue-4)
# Channels
* [How do you create a channel for goroutine?](#Usage)
* [How do you create a buffered channel for goroutine?](#Usage)
* [Thinks about the difference between a unbuffered channel (default) and a buffered channel?](#Usage)
* [How do you create a channel directly?](#Usage)
 

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
Ouput
```
=== RUN TestPrint1
0 1 2 3 4 5 6 7 8 9 A B C D E F G H I J --- PASS: TestPrint1 (0.00s)
=== RUN TestGoPrint1
--- PASS: TestGoPrint1 (0.00s)
PASS
```
#### Issue-1
If you can't see output at test case 2, because the processing work being done before goroutine display the output. 
To fix this issue, you need to add a line code.
```
func TestGoPrint1(t *testing.T) {
    goPrint1()
    time.Sleep(1 * time.Millisecond)
}
```
Output
```
=== RUN TestPrint1
0 1 2 3 4 5 6 7 8 9 A B C D E F G H I J --- PASS: TestPrint1 (0.00s)
=== RUN TestGoPrint1
0 1 2 3 4 5 6 7 8 9 A B C D E F G H I J --- PASS: TestGoPrint1 (0.00s)
PASS
```
#### Issue-2
As you see, both run the same  way with the same output.  
The reasonyou get the same results is that printNumbers1 and printLetters1 ran so quickly, it made no difference whether or not the functions were running independently.  
To simulate processing work, you’ll add a time delay using the Sleep function in the time package, and re-create the two functions as printNumbers2 and printLetters2 in **goroutine.go**.
```
func printNumbers2() {
    for i := 0; i < 10; i++ {
        time.Sleep(1 * time.Microsecond
        fmt.Printf("%d ", i)
    }
}
func printLetters2() {
    for i := 'A'; i < 'A'+10; i++ {
        time.Sleep(1 * time.Microsecond)
        fmt.Printf("%c ", i)
    }
}
func goPrint2() {
    go printNumbers2()
    go printLetters2()
}
```
Output
```
=== RUN TestPrint1
0 1 2 3 4 5 6 7 8 9 A B C D E F G H I J --- PASS: TestPrint1 (0.00s)
=== RUN TestGoPrint1
0 1 2 3 4 5 6 7 8 9 A B C D E F G H I J --- PASS: TestGoPrint1 (0.00s)
=== RUN TestGoPrint2
A 0 B 1 C D 2 E 3 F 4 G H 5 I 6 J 7 8 9 --- PASS: TestGoPrint2 (0.00s)
PASS
```
#### Issue-3
If you run this code again, the last line produces a differnt result. In fact, **printNumbers2** and **printLetters2** run independently and fight to print to the screen. Running repeatedly will produce different results each time. If you’re using a Go version prior to Go 1.5, you might get the same results each time. Why?  

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
#### Work-flow-using-WaitGroup
To get more details about code. Refer: [here](https://github.com/huavanthong/MasterGolang/blob/feature/chapter9/01_GettingStarted/book-go-web-application/Chapter_9_Leveraging_Go_Concurrency/9.2.3_Waiting_for_goroutine/waiting.go)
> You set to waitGroup with counter 2, and you make 2 goroutine in your program. 
> => You need to wait the counter becomes 0. If waiting's done, you finish process. 
> Otherwise, you continue to wait it.

#### Issue-4
If you forget to decrement the counter. And you run this program, we will make error below
```
0 A 1 B 2 C 3 D 4 E 5 F 6 G 7 H 8 I 9 J fatal error: all goroutines are
asleep - deadlock!
```
## Channels
* In the previous example, you saw how the **go** keyword can be used to convert normal functions into goroutines and execute them indenpendently. [here](#Using-goroutines)
* In the last subsection, you also saw how to use WaitGroups to synchronize between independently running goroutines. [here](#Waiting-for-goroutines)

In this section, you’ll learn how goroutines can communicate with each other using channels. More details: [here](https://www.meisternote.com/app/note/F23PWHgBDY5M/9-3-channels)
### Usage
To allocates a channel of integer
```
ch := make(chan int)
```
Channels are, by default, unbuffered.  If an optional integer parameter is provided, a buffered channel of the given size is allocated instead. This creates a buffered channel of integers with the size 10:
```
ch := make(chan int, 10)
```
The syntax for putting things into a channel is quickly recognizable, visually. This puts an integer 1 into the channel ch:
```
ch <- 1
```
Taking out the value from a channel is equally recognizable. This removes the value from the channel and assigns it to the variable i:
```
i := <- ch
```
Channels can be directional. By default, channels work both ways (bidirectional) and values can be sent to or received from it. But channels can be restricted to send-only or receive-only. This allocates a send-only channel of strings:
```
ch := make(chan <- string)
```
This allocates a receive-only channel of strings:
```
ch := make(<-chan string)
```
### Synchronization-with-channels
Refer: [here](https://github.com/huavanthong/MasterGolang/tree/feature/chapter9/01_GettingStarted/book-go-web-application/Chapter_9_Leveraging_Go_Concurrency/9.3.1_Synchronization_with_channels)
