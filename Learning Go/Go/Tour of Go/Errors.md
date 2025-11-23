Go expresses errors with `error` values

The `error` type is a built-in interface similar to `fmt.Stringer`:

```Go
type error interface {
    Error() string
}
```

(As with `fmt.Stringer`, the `fmt` package looks for the `error` interface when printing values.)

Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.

```Go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)

```

## Example
```Go
import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	//Go structs defined like this are considered to be adresable so we can call
	// &MyError{} directly
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```

## Exercise, add error handling to previous Newton formula


**Note:** A call to `fmt.Sprint(e)` inside the `Error` method will send the program into an infinite loop. You can avoid this by converting `e` first: `fmt.Sprint(float64(e))`. Why?

>This is because if we don't convert `num` which in this scenario is a type that implements error it will call back Error on it and then end on an infinite loop. We need to convert it to `float64` so it no longer implements error interface and it does not create a stack exception

```Go
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64


func (num *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(*num))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		//We need to separate this into an assignment  then we do &err
		//This is because temporary values such as a type conversion
		// like ErrNegativeSqrt(x) are not addresable
		// that means something like &ErrNegativeSqrt(x) will fail, structs work
		err := ErrNegativeSqrt(x)
		return 0, &err
	}
	
	z := 1.0
	old_z := 1.0
	for a:=1; a <= 10; a++ {
		old_z = z
		z -= (z*z - x) / (2*z)
		fmt.Printf("z value %.10f\n", z)
		
		if math.Abs(z - old_z) < 1e-6 {
			//here we break since we don't see change on the 6 decimal mark
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

```



- `num` is a pointer (`*ErrNegativeSqrt`)
- `*num` dereferences it to the actual value (`ErrNegativeSqrt`)
- `float64(*num)` converts your custom type into the built-in type
- Only needed because your custom type **is a float, not a struct**