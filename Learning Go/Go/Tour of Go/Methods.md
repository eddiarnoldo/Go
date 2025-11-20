Go does not have classes, however you can define methods on types.
A method is a special function with a `receiver` argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

```Go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

```


Another example:

```Go
func (v Vertex) Sum() float64 {
	return v.X + v.Y
}
```

Methods are functions with a receiver argument, here is the same Abs method but as a function

```Go
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

You can also declare method on non-struct types

```Go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

```

> Here we can see `MyFloat` is from type float64, but we can still add a method for MyFloat

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

## Pointer receivers
You can declare methods with pointer receivers.

This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)

For example, the `Scale` method here is defined on `*Vertex`.

Methods with pointer receivers can modify the value to which the receiver points (as `Scale` does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the `*` from the declaration of the `Scale` function on line 16 and observe how the program's behavior changes.

With a value receiver, the `Scale` method operates on a copy of the original `Vertex` value. (This is the same behavior as for any other function argument.) The `Scale` method must have a pointer receiver to change the `Vertex` value declared in the `main` function.

```Go
package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Sum() float64 {
	return v.X + v.Y
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Sum())
}

```

## Methods and pointer indirection
```Go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

```Go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

```

Here `v.Scale(5)` works since Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

This does not work the same way for functions i.e if we did `ScaleFunc(v)` it won't compile

> The opposite happens if the function receives a value and not a pointer argument

```Go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!


var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK. this is interpreted as (*p).Abs()
```

```Go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

## Choosing a value or a pointer receiver
There are 2 reasons on why to use a pointer receiver

- So the method can modify the value that its receiver points to
- To avoid copying the value on each method call. This can be more efficient if the receiver is a large struct.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.