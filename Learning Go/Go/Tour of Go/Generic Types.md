
In addition to generic functions, Go also supports generic types. A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.

This example demonstrates a simple type declaration for a singly-linked list holding any type of value.

As an exercise, add some functionality to this list implementation.

```Go
package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.

//Here the generyc type goes within [T any]
type List[T any] struct {
	next *List[T]
	val  T
}

type Person struct { Name string Age int }

func (l *List[T]) Add(elem T) List[T] {
	//We use & since we need to give a pointer to the List not the actual value
	next := List[T]{val: elem}
	l.next = &next
	return next
}

// Get returns the value at this node
func (l *List[T]) Get() T {
	return l.val
}

// Next returns the next node
func (l *List[T]) Next() *List[T] {
	return l.next
}

func main() {
	list := List[string]{val: "Hello"}
	next := list.Add("World")
	exclamation := next.Add("!")
	
	linked := []List[string]{list, next, exclamation}
	
	//or
		
	//linked := make([]List[string], 0)
	//linked = append(linked, list)
	//linked = append(linked, next)
	//linked = append(linked, exclamation)
	
	//Message from generic type slice
	for _, elem := range linked {
		fmt.Printf("%v ", elem.Get())
	}
	
	personList := List[Person]{val: Person{"Eddy", 34}}
	fmt.Println(personList.Get().Name)
}

```
