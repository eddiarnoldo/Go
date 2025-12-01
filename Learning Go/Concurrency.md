# Goroutines

A _goroutine_ is a lightweight thread managed by the Go runtime.

```Go
go f(x, y, z)
```

starts a new goroutine running

```Go
f(x, y, z)
```

The evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of `f` happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The [`sync`](https://go.dev/pkg/sync/) package provides useful primitives, although you won't need them much in Go as there are other primitives.

```Go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() { 
	go say("world") // Start goroutine 
	say("hello") // This blocks for ~500ms (5 iterations × 100ms) 
	// By the time say("hello") finishes, say("world") has already run }

```

> If you comment `say("hello")` nothing happens since the main method finishes and it stops execution!

# Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

```Go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:
```Go
ch := make(chan int)
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

```Go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y:= <-c, <-c // receive from c

	fmt.Println(x, y,z, x+y)
	
	//If we did x, y, z := <-c, <-c, <-c 
	//The code will panic and say "fatal error: all goroutines are asleep - deadlock!"
}

```

# Buffered channels
Channels can be _buffered_. Provide the buffer length as the second argument to `make` to initialize a buffered channel:

```Go
ch := make(chan int, 100)
```

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.

## Notes
Buffered channels means you can send up to 100 values to the channel without anyone receiving them yet, but if you sent 101 it will block until someone frees up some of the buffer.

## How unbuffered channels work:

An unbuffered channel has **no storage capacity**. This means:

- **Every send blocks until there's a receiver ready**
- **Every receive blocks until there's a sender ready**

It's **synchronous communication** - the sender and receiver must "meet" at the same time.

## Example:

go

```go
ch := make(chan int)  // unbuffered

ch <- 42  // BLOCKS immediately! 
          // No buffer to store the value
          // Waits for someone to receive
```

This would deadlock if done in the main goroutine alone.

## Working example with goroutine:

go

```go
ch := make(chan int)  // unbuffered

go func() {
    ch <- 42  // Sends, then waits for receiver
}()

x := <-ch  // Receives - now the send completes
fmt.Println(x)  // 42
```

# Range and Close
A sender can `close` a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```Go
v, ok := <-ch
```

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.

**Note:** Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

**Another note:** Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.

```Go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//This is what closes the channel
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c)+1, c)
	for i := range c {
		fmt.Println(i)
	}
	
	//We do this here since if we did it earlier in before the loop we will consume one of the values
	_, ok := <-c
	fmt.Println(ok)
}

```

# Select

The `select` statement lets a goroutine wait on multiple communication operations.

A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

```Go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

```

## Step-by-step execution:

1. **Anonymous goroutine starts**, waits at `<-c` (blocked, waiting for a value)
2. **`fibonacci()` runs** in main goroutine, hits the select:

go

```go
   select {
       case c <- x:        // CAN send (goroutine is waiting)
       case <-quit:        // CAN'T receive (nothing sent yet)
   }
```

- The first case executes, sends `0` to `c`

3. **Anonymous goroutine receives** `0`, prints it, loops back to `<-c` (blocked again)
4. **`fibonacci()` loops**, tries `c <- x` again with `x=1`
    - Blocks until the goroutine is ready to receive
    - Goroutine receives, prints `1`
5. **This repeats 10 times** (i goes from 0 to 9)
6. **After iteration 9**, the anonymous goroutine finishes the loop and executes:

go

```go
   quit <- 0
```

7. **Now in `fibonacci()`**, the select has:

go

```go
   select {
       case c <- x:     // Blocked (no receiver waiting anymore)
       case <-quit:     // CAN receive (quit has a value!)
   }
```

- The second case executes, prints "quit", and returns

## The answer to your question:

`case c <- x:` stops being chosen because **there's no receiver ready** anymore. The select then chooses the `case <-quit:` instead, which has a value waiting.