# Slices

A slice is like an array, except that its size changes dynamically as you add and remove items. The slice is one of Go's most useful types.

```go
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)
 

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// This is a lowercase function name which makes it private basically
// accessible only to code in its own package (in other words, it's not exported)
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v well met!",
	}

	return formats[rand.Intn(len(formats))]
}
```

As we can see in the code above a slice is defined by 

```go
formats := []string(
	"1",
	"2",
	"3"	
)
```

