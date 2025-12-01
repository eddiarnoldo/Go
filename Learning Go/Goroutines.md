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

