An `interface` is defined as a set of method signatures

```Go
type Abser interface {
	Abs() float64
}
```

```Go

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```


Here we can see MyFloat implements Abs, and Vertex also implements it however the method uses a pointer receiver  that's why `a=v` fails since `*Vertex` is the one implementing it

If we change the method Abs to use a value `receiver` it will work.

```Go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

^ if we do this change both `a = &v` and `a = v` work since the method using the value receiver makes both `Vertex` and `*Vertex` have it.

When you have a pointer (`*T`) and the method has a value receiver, Go can automatically dereference it.

## Interfaces are implemented implicitly

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.


```Go
package main

import "fmt"

type I interface {
	M()
	hello() string
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	fmt.Println(t.S)
}

func (t *T) hello() string {
	return "hello method"
}

func main() {
	// here we had do to &T{"hello"} Since *T
	// is the one that implements the interface
	var i I = &T{"hello"} 
	i.M()
	fmt.Println(i.hello())
}
```

## Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)

```Go

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```


Prints
```
(&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
3.141592653589793
```

> Calling a method on an interface value executes the method of the same name on its underlying type.

## Interface values with nil underlying values
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method `M` in this example.)

```Go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

Prints

```
(<nil>, *main.T)
<nil>
(&{hello}, *main.T)
hello
```


## Nil interface values
A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which _concrete_ method to call.

```Go
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

This will fail since this is a nil interface 
```
(<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4992b9]

goroutine 1 [running]:
main.main()
	/tmp/sandbox2995677200/prog.go:12 +0x19
```

## Empty interface

The interface type that specifies zero methods is known as the _empty interface_:

```Go
interface{}
```

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

```Go
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

Prints
```
(<nil>, <nil>)
(42, int)
(hello, string)
```

Everything implements the Empty interface!!

## Type assertions

A type assertion provides access to an interface value's underlying concrete value.

```Go
t  := i.(T)
```

If `i` does not hold a `T`, the statement will trigger a panic.

To _test_ whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

```Go
t, ok := i.(T)
```

```Go
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

> This is very similar to when we check if a map has an entry for a given key
>`elem, ok = m[key]`


## Type switches
A _type switch_ is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```Go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

The declaration in a type switch has the same syntax as a type assertion `i.(T)`, but the specific type `T` is replaced with the keyword `type`.


```Go
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("This is a bool with value %v \n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(24.56)
}
```

Here we can see the do function receives a empty interface parameter and then uses a type switch to validate the types.

## Stringers
One of the most ubiquitous interfaces is [`Stringer`](https://go.dev/pkg/fmt/#Stringer) defined by the [`fmt`](https://go.dev/pkg/fmt/) package.

```Go
type Stringer interface {
    String() string
}
```

A `Stringer` is a type that can describe itself as a string. The `fmt` package (and many others) look for this interface to print values.

```Go

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	
	b := Game{"Leyend of Zelda Ocarina of Time", 10.0}
	fmt.Println(b)
}

// I created this!!
type Game struct {
	Title string
	Rating float64
}

func (g Game) String() string {
	return fmt.Sprintf("This game title is %v, and it has a rating of %f", g.Title, g.Rating)
}
```

