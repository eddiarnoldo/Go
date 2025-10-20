### Declaring variables
`var` can be used to declare variables in Go, it can also be used to declare multiple at a time and can also omit the type if they are the same as long as we declare it till the end.

The can also be declared at the package or function level

```go
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

```

### Variable initializers
`var` declaration can include initializers if this happens the type can be omitted

```go
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

### Short variable declaration
**This just works inside functions**, we can use `:=` to do assignments, outside of the functions we need to use `var`, `func` constructs `:=` is not available.

```go
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```
