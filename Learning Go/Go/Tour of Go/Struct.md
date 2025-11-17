A Struct is a collection of fields

```Go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

You can access the fields from a `struct` using a dot

```Go
func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```

Struct fields can be accessed through a struct pointer.

```Go
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 //this would have been (*p).X but Go allows a simple call with p.X
	fmt.Println(v)
}
```

Declare structs samples

```Go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
```