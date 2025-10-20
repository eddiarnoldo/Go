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
