The `range` form of the `for` iterates over a slice or map

```Go
import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

## Range continued
You can use `_` if you want to skip one the values, if you just want the index you can omit the `_`

```Go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

`
### Exercise
```Go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Create the outer slice of length dy
	image := make([][]uint8, dy)
	
	// For each row, create an inner slice of length dx
	for y := 0; y < dy; y++ {
		image[y] = make([]uint8, dx)
		
		// Fill in the pixel values
		for x := 0; x < dx; x++ {
			// Try different formulas here!
			image[y][x] = uint8((x + y) / 2)
			// image[y][x] = uint8(x * y)
			// image[y][x] = uint8(x ^ y)  // XOR operation
		}
	}
	
	return image
}

func main() {
	pic.Show(Pic)
}
```

https://go.dev/tour/moretypes/18