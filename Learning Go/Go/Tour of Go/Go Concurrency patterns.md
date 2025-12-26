Link https://www.youtube.com/watch?v=f6kdp27TYZs

# Concurrency is not parallelism

- Concurrency is not parallelism, although it enables parallelism
- If you have one processor your program can still be concurrent but cannot be parallel


Concurrency should be:
- Easy to understand
- Easy to use
- Easy to reason about
- You don't need to be an expert

---
> Go is about channels!!

## New thing learned aside from Go
Adding `&` at the end of a shell command runs it in the **background**.

For example:

bash

```bash
long_running_process &
```

This launches `long_running_process` and immediately returns control to your terminal, so you can run other commands.

**Key behaviors:**

- The process runs in the background while you can use the shell normally
- You get a job number: `[1] 12345` (where 1 is the job number, 12345 is the PID)
- You can still see output from the background process mixed with your terminal
- Typing `fg` brings the most recent background job back to the foreground
- Typing `jobs` shows all running background jobs
- Typing `bg` resumes a suspended job in the background

Back to go 


## Goroutines
- It's an independently executing function launched by a `go` statement
- It has it's own call stack, which grows and shrinks as required
- It's very cheap thousands or even hundred of thousands of them
- It's not a thread
- There might be single thread with thousands of goroutines


## Communication
Use channels! 
```Go
var c = chan int
c = make(chan int)

// or

d := make(chan int)
```

```Go
c <- 1 //send 1 to the channel

value = <-c
```

## Synchronization
When a main function executes `<-c` it will wait for a value to be sent

Similarly, when the boring function executes `c <- value`, it waits for a receiver to be ready.

> A sender and a receiver must both be ready to play their part on the communication, Otherwise we wait until they are

### Buffered channels
- Buffering removes synchronization
- Buffering makes them more like Erlang Mailboxes
- Buffered channels can be important for some problems but they are more subtle to reason about


# The Go Approach

DON'T COMMUNICATE BY SHARING MEMORY, SHARE MEMORY BY COMMUNICATING

> In other words you don't have a blob of memory and add a bunch of mutexes and condition variables around it to protect it from parallel access. Instead you use the channel to pass the data back and forth between the goroutines and make your concurrent program operate that way.

## Patterns

### Generators

Functions that return a channel 

```Go
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	joe := boring("Joe")
	sam := boring("Sam")
	
	for i:=0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-sam)
	}
	
	fmt.Println("I'm leaving")
}

func boring(msg string) chan string {
	c := make(chan string)
	go func(){
		for i := 0; ;i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.IntN(1e3)) * time.Millisecond)
		}
	}()
	return c //Return the channel to the caller
}

```

### Channels as a handle on a service
Boring function returns a channel that lets us communicate with the boring service it provides

```Go
func main(){
	joe := boring("Joe")
	sam := boring("Sam")
	
	for i:=0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-sam)
	}
	
	fmt.Println("I'm leaving")
}
```

```bash
Joe 0
Sam 0
Joe 1
Sam 1
Joe 2
Sam 2
Joe 3
Sam 3
Joe 4
Sam 4
I'm leaving
```

Because of the synchronization nature of the channel `sam` is blocked until Joe receives data from the channel.

This could be bad if Sam is more talkative

### Multiplexing!

These programs make Joe and Ann count in lockstep.
We can instead use a fan-in function to let whosoever is ready talk.


```Go
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func(){ for {c <- <-input1} }()
	go func(){ for {c <- <-input2} }()
	return c
}

func main(){
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Bye")
}
```

> Go has bidirectional channels, receive only channels and send only channels 

```Go
ch := make(chan string) // Bidirectional
var recvOnly <-chan string = ch // Now recvOnly is receive-only
var sendOnly chan<- string = ch // Now sendOnly is send-only
```


![[Pasted image 20251210214855.png]]

### Restoring sequencing
Send a channel on a channel, making a goroutine wait its turn

Receives all messages, then enable them again by sending on a private channel

```Go
type Message struct {
	str string
	wait chan bool
}
```

```Go
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Eddy"))
	
	for i:=0; i < 10; i++ {
		msg1 := <-c; fmt.Println(msg1.str)
		msg2 := <-c; fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	
	fmt.Println("I'm leaving")
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func(){ for {c <- <-input1} }()
	go func(){ for {c <- <-input2} }()
	return c
}

type Message struct {
	str string
	wait chan bool
}

func boring(msg string) chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func(){
		for i := 0; ;i++ {
			//Here we create a message with the waitforit inside of it
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.IntN(1e3)) * time.Millisecond)
			<-waitForIt //This locks until someone confirms as received
		}
	}()
	return c //Return the channel to the caller
}
```

Prints 
```bash
Eddy 0
Joe 0
Eddy 1
Joe 1
Eddy 2
Joe 2
Eddy 3
Joe 3
Eddy 4
Joe 4
Eddy 5
Joe 5
Eddy 6
Joe 6
Eddy 7
Joe 7
Eddy 8
Joe 8
Eddy 9
Joe 9
I'm leaving
```

Now they are back in sync!!

## Select

Is a control structure similar to switch, used to control the behavior of your program based on what communications are able to proceed at any moment.

```Go
select {
	case v1 := <-c1:
		fmt.Printf("Received %v from c1\n", v1)
	case v2 := <-c2:
		fmt.Printf("Received %v from c2\n", v2)
	case c3 <- 23:
		fmt.Printf("sent %v to c3\n", 23)
	default:
		fmt.Printf("no one was ready to communicate\n")	
}
```

> We can rewrite our `fanin` function with select

```Go
func fanIn(input1, input2 <- chan string) <-chan string {
	c := make(chan string)
	go func(){
		for {
			select {
				case s := <-input1: c <- s
				case s := <-input2: c <- s
			}
		}
	}()
	return c
}
```

### Timeout using select
The `time.After` function returns a channel that blocks for the specified duration. After the interval, the channel delivers the current time, once

```Go
func main() {
	c := boring("Joe")
	for {
		select {
			case s := <-c:
				fmt.Println(s)
			case <-time.After(1 * time.Second):
				fmt.Println("You're to slow. Bye")
				return	
		}
	}
}
```

> Above is timing after each message 

or we could do 

```Go
func main() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
			case s := <-c:
				fmt.Println(s)
			case <-timeout:
				fmt.Println("You're to slow. Bye")
				return	
		}
	}
}
```

> This one times the whole conversation
## Quit channel!

We can send a channel that we can use as a quit to the go routine running the code to Stop when we're tired of listeting

```Go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	c := boring("Eddy", quit)
	for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c)}
	quit <- true
	
}


func boring(str string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func(){
		for i:= 0 ; ; i++ {
			select {
				case c <- fmt.Sprintf("%s, %d", str, i):
					//do nothing
				case <-quit:
					fmt.Println("Told me to quit Bye!")
					return
			}
		}
	}()
	return c
}
```

>  There is a problem with above example there might be some cleanup, so we can change it to do some cleanup and use a bidirectional channel so main waits for the cleanup \

### Receive on quit channel improvement (wait for cleanup)

```Go
func main() {
	quit := make(chan string)
	c := boring("Eddy", quit)
	for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c)}
	quit <- "Bye!"
	fmt.Printf("Eddy says: %q\n", <-quit)
	
}


func boring(str string, quit chan string) <-chan string {
	c := make(chan string)
	go func(){
		for i:= 0 ; ; i++ {
			select {
				case c <- fmt.Sprintf("%s, %d", str, i):
					//do nothing
				case <-quit:
					cleanup()
					quit <- "See ya"
					return
			}
		}
	}()
	return c
}

func cleanup() {
	fmt.Println("Cleaning..")
}
```

### Daisy-chain
```Go
func f(left, right chan int) {
	left <- 1 + <- right
}

func main(){
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	
	go func(c chan int){ c <- 1}(right)
	fmt.Println(<-leftmost)
}
```

>Above's code basically creates 10000 go routines that play chinese whispers

gorutine 1 -> 2 -> 3 .... 10000


## System Software

Go was designed for writing systems software, let's see how the concurrency features come into play

### Example: Google Search

```
Q: What does Google search do?
A: Given a query, return a page of search results (and some ads)

Q: How do we get the search results?
A: Send the query to Web Searc, Image search, YouTube, Maps, etc
then mix results
```

How do we implement this?



--- Pending to add more notes see goo. google.go file

## Don't overdo it
Use the right tool for the right job, don't over use go routines etc if they are not needed.

## Conclusion

