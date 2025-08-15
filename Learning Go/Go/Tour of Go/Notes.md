## Packages

Programs are made of packages
> By convention package name is the same as the last element of the import path 
> For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`.


```go
package main

import (
	"fmt"
	"math/rand"
)
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

## Imports

### Factored

```go
import (
	"fmt"
	"math/rand"
)
```

### Non factored

```go
import "fmt"
import "math/rand"


func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

> `%g` is  is a **floating-point formatting verb**.


## Exported names

Names that start with capital letters are exported if they do not start with capital letter they are not exported

```go

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi) //This works
	fmt.Println(math.Pi) //This fails
}
```

