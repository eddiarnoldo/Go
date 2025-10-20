A map maps keys to values.

The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.

The `make` function returns a map of the given type, initialized and ready for use.


```Go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

```

## Map literals
```Go
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```

> If the top level is just a type name you can omit it from the elements of the literal

```Go
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

```

### Insert

```Go
m[key] = elem
```

### Retrieve
```Go
elem = m[key]
```

### Test a key
```Go
elem, ok = m[key]

elem, ok := m[key] //if elem or ok have not been declared
```

## Exercise Word count

```Go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)
	
	for _, val := range words {
		wordCount[val]++
	}
	
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
```