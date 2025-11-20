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

