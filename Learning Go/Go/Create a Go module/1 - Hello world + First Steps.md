## How to init a project?

```
go mod init {module name}
```

This will create the basic Go project structure for a module
- `go.mod`
- `go.sum`

## Structure Hello World example

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## FMT module


Contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.


## How to run a program

`go run .`

