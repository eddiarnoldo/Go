
Functions can take 0 or n arguments, types come after the variable name

```go
func add(x int, y int) int {
	return x + y
}
```

### Omit types
When two or more consecutive named function parameters share the type you can omit the type except for the last one.

```go
func add(x, y int) int {}
```

### Multiple results

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

### Named return values
Go functions returns values may be named if they are named they are threated as variables declared at the top of the function,, When we do a return with not arguments this is called a `naked` return and in this scenario Go function returns those named return values

```go
func split(sum int) (x , y int) {
	x = sum * 4 / 9
	y = sum - x
	return // <- this is the naked return 
}
```