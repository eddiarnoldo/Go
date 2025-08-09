Visit `pkg.go.dev` to download published modules so you can use them in you project.

### How to fetch a dependency?

#### Manual

`go get rsc.io/quote`

#### Automatic

`go mod tidy`
  
- Automatically download missing dependencies like rsc.io/quote
- Add them to your go.mod and go.sum
- Remove any unused modules
    

## Modules


A function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an `exported name`.

```go
package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi %v, Welcome", name)
	return message
}
```
  

In Go, the `:=` operator is a shortcut for declaring and initializing a variable in one line
  

```go
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```

vs

```go
message := fmt.Sprintf("Hi, %v. Welcome!", name)
```

## Use replace for local modules still no published

Normally you would use modules that are published but in this scenario we are creating a module `examples.com/greetings` which is local.

This means if we simply use `go mod tidy` it won't be able to find the module we are tying to use.

In order to achieve this locally we need to use

```bash
go mod edit -replace example.com/greetings=../greetings
```

This way `go` will modify the `go.mod` and a replace so we can tell our project to look locally.

After this we can use `go mod tidy` and it will work 

```go
module example.com/hello

go 1.24.6  

replace example.com/greetings => ../greetings  

require example.com/greetings v0.0.0-00010101000000-000000000000
```

