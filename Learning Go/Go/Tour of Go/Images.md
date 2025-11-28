[Package image](https://go.dev/pkg/image/#Image) defines the `Image` interface:

```Go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

**Note**: the `Rectangle` return value of the `Bounds` method is actually an [`image.Rectangle`](https://go.dev/pkg/image/#Rectangle), as the declaration is inside package `image`.

(See [the documentation](https://go.dev/pkg/image/#Image) for all the details.)

The `color.Color` and `color.Model` types are also interfaces, but we'll ignore that by using the predefined implementations `color.RGBA` and `color.RGBAModel`. These interfaces and types are specified by the [image/color package](https://go.dev/pkg/image/color/).

```Go
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

```

## Exercise Images
Remember the [picture generator](https://go.dev/tour/moretypes/18) you wrote earlier? Let's write another one, but this time it will return an implementation of `image.Image` instead of a slice of data.

Define your own `Image` type, implement [the necessary methods](https://go.dev/pkg/image/#Image), and call `pic.ShowImage`.

`Bounds` should return a `image.Rectangle`, like `image.Rect(0, 0, w, h)`.

`ColorModel` should return `color.RGBAModel`.

`At` should return a color; the value `v` in the last picture generator corresponds to `color.RGBA{v, v, 255, 255}` in this one.

```Go
package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	Width, Height int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	//v := uint8(x * y) // or any formula you like
	// color.RGBA{uint8(v), uint8(v), 255, 255}
	return color.RGBA{uint8(x), uint8(y), 150, 255}
}

func main() {
	m := Image{Width: 256, Height: 256}
	pic.ShowImage(m)
}
```

![[Pasted image 20251122235931.png]]

For reference:

```Go
type RGBA struct { R, G, B, A uint8 }
```

This creates a color where:

- **Red = v** (varies based on pixel position)
- **Green = v** (varies based on pixel position)
- **Blue = 255** (always maximum blue)
- **Alpha = 255** (always fully opaque)