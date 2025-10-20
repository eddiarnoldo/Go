

```go
sum := 0

for a := 0; a < 10; a++ {
	sum += a
}

fmt.Println(sum)
```

Go has only one looping construct, the `for` loop.

The basic `for` loop has three components separated by semicolons:

- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration


```go
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```

### For is  `Go` While

> You can drop the `;`

```go
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

### For (forever)

```go
for {
	}
```

