A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```Go
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
```

So `defer` is like having multiple tiny `finally` blocks that you can sprinkle throughout your function, and they execute in reverse order. Much more flexible than Java's single `finally`!

`defer` go into a stack

```Go

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```

```
counting
done
9
8
7
6
5
4
3
2
1
0
```

