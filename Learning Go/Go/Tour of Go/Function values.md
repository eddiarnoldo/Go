Functions are values to, they can be passed around just like other values.

```Go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

We need to use `fn` on the function parameters to receive a function, then we define the signature of this function.

## Function closures
Functions may be closures, A closure is a function value that references variables from outside its body.  The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

```Go
package main

import "fmt"

//Here we can see that adder returns a func(int) int, and sum becomes it's bound 
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

```bash
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90
```

## Exercise: Fibonacci closure

```Go
package main
import "fmt"

func fibonacci() func() int {
	series := []int{0, 1}
	cur := 0
	return func() int {
		if cur >= len(series) {
			nextVal := series[cur-2] + series[cur-1]
			series = append(series, nextVal)
		}
		result := series[cur]
		cur++
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```