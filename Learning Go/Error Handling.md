We can use package `errors` to be able to return errors in our code.

```go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
```

Here as we can see if the name is empty we return a new error, notice that our function now returns `two` values. In `Go` functions can return multiple values.

## Handling error

Now that our module can return errors we need to handle them

```go
package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)
  
func main() {
	log.SetPrefix("greetings: ")
	//log.SetFlags(0)
	message, err := greetings.Hello("")
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
```

Here we can see that we are receiving the 2 values of our function call and we act based on if there is an error that exists. 

`log.Fatal` logs the error and stops the execution of the program

With log.SetFlags(0)
```bash
greetings: 2025/08/09 21:22:20 empty name
exit status 1
```

Without 
