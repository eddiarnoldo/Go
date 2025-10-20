Same as the `for` the `if` does not need the `()`

```go
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
```

> `if` also supports short statements

```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```

` v := math.Pow(x, n);` this only exists in the context of the `if` also in the `else` body

#### Exercise implement Newton formula for sqrt

```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	old_z := 1.0
	for a:=1; a <= 10; a++ {
		old_z = z
		z -= (z*z - x) / (2*z)
		fmt.Printf("z value %.10f\n", z)
		
		if math.Abs(z - old_z) < 1e-6 {
			//here we break since we don't see change on the 6 decimal mark
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
```

##### Output:

```bash
z value 1.5000000000
z value 1.4166666667
z value 1.4142156863
z value 1.4142135624
z value 1.4142135624
1.4142135623730951
```

