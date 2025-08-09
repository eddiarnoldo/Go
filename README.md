# Go
This repo will be used to learn to program on `Go`

## How to init a project?

```
go mod init {module name}
```

## Structure Hello World example
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## FMT

 Contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.


## How to run a program

`go run .`

## Dependencies?

Visit `pkg.go.dev` to download published modules so you can use them in you project.

### How to fetch a dependency?
#### Manual
`go get rsc.io/quote`

#### Automatic
`go mod tidy`

- Automatically download missing dependencies like rsc.io/quote
- Add them to your go.mod and go.sum
- Remove any unused modules