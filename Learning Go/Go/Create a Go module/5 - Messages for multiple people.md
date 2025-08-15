Now we want to add the ability to our module to create greeting for multiple people, however there is a catch if we simply modified our Hello function the users of it will have issues due to it's signature change.

That's why we will create a new function that can receive the new set of parameters.

```go

func Hellos(names []string) (map[string]string, error) {
	//Create a map to associate messages to persons
	messages := make(map[string]string)

	//This syntax is because `range names` returns to parameters
	//	- The index of the current item in the loop
	//  - A copy of the item's value
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}

		messages[name] = message
	}

	return messages, nil
}
```

## Maps in go
Maps are created with the following syntax:
```go
// `make(map[_key-type_]_value-type_)`
messages := make(map[string]string)
```

### Iterating a slice
```go
for _, name := range names 
```

The `range` function returns 2 parameters back from it's call, the index and the value on the slice, however since we don't need really have a use case for the index we can use the `_` called the [blank identifier](https://go.dev/doc/effective_go.html#blank)

## Modify our main func

Now that we have a new `Hellos` function we can update our main code to call this function from our module sending a slice of users

```go
func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}
```