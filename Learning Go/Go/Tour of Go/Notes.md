## Packages

Programs are made of packages
> By convention package name is the same as the last element of the import path 
> For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`.


```go
package main

import (
	"fmt"
	"math/rand"
)
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

## Imports

### Factored

```go
import (
	"fmt"
	"math/rand"
)
```

### Non factored

```go
import "fmt"
import "math/rand"


func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

> `%g` is  is a **floating-point formatting verb**.


## Exported names

Names that start with capital letters are exported if they do not start with capital letter they are not exported

```go

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) //This works
	fmt.Println(math.Pi) //This fails
}
```

## Functions
Functions can take 0 or n arguments, types come after the variable name

```go
func add(x int, y int) int {
	return x + y
}
```

### Omit types
When two or more consecutive named function parameters share the type you can omit the type except for the last one.

```go
func add(x, y int) int {}
```

### Multiple results

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

### Named return values
Go functions returns values may be named if they are named they are threated as variables declared at the top of the function,, When we do a return with not arguments this is called a `naked` return and in this scenario Go function returns those named return values

```go
func split(sum int) (x , y int) {
	x = sum * 4 / 9
	y = sum - x
	return // <- this is the naked return 
}
```

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

### Basic types

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

```go
import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	r      rune       = '☺'
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println(r)       // prints 9786
	fmt.Printf("%T %v %c", r, r, r) // prints int32 9786 ☺
}
```

> The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements


### Default values

Variables declared without a initial value are given their `zero` value

- `0` for numeric types,
- `false` for the boolean type, and
- `""` (the empty string) for strings.


### Type conversions

The expression `T(v)` converts the value `v` to the type `T`.

```go
import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```


> Notice that if we change `math.Sqrt(float64(x*x + y*y))` to `math.Sqrt(9)` Go will infer the type since 9 is an `Untyped` constant so go transforms it to `float64` but if you do `math.Sqrt(x*x + y*y)` it will fail since both x and y multiplications produce an int



### Type inference

This is when we use the `:=` construct

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```


### Constants

Constants do not necessarelly need a type in Go

- Go constants are **compile-time values**, not runtime variables.
- Untyped constants give **flexibility** — they can be used wherever a compatible type is expected.

They use the word `const` , can't use the `:=` syntax

```go
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```


#### Numeric constants
Numeric constants are high-precision _values_.
An untyped constant takes the type needed by its context.
Try printing `needInt(Big)` too.

```go
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```
