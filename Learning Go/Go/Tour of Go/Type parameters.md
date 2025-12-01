Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets, before the function's arguments.

```Go
func Index[T comparable](s []T, x T) int
```

> This declaration means that `s` is a slice of any type `T` that fulfills the built-in constraint `comparable`. `x` is also a value of the same type.

`comparable` is a useful constraint that makes it possible to use the `==` and `!=` operators on values of the type. In this example, we use it to compare a value to all slice elements until a match is found. This `Index` function works for any type that supports comparison.

```Go
package main

import "fmt"

type Person struct {
	name, lastname string
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
	
	// Structs are also comparable as long as their fields are comparable
	sstruct := []Person{
		{"Eduardo", "Quintana"},
		{"Rocio", "Andujo"},
		{"Leonel", "Quintana-Andujo"},
	}
	
	fmt.Println(Index(sstruct, Person{"Leonel", "Quintana-Andujo"}))
}
```