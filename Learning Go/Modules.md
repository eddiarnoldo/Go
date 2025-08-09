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

