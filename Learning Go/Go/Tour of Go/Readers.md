The `io` package specifies the `io.Reader` interface, which represents the read end of a stream of data.

The Go standard library contains [many implementations](https://cs.opensource.google/search?q=Read%5C\(%5Cw%2B%5Cs%5C%5B%5C%5Dbyte%5C\)&ss=go%2Fgo) of this interface, including files, network connections, compressors, ciphers, and others.

The `io.Reader` interface has a `Read` method:
	
```Go
func (T) Read(b []byte) (n int, err error)
```


`Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.

The example code creates a [`strings.Reader`](https://go.dev/pkg/strings/#Reader) and consumes its output 8 bytes at a time


```Go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```

## Output
```bash
n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
b[:n] = "Hello, R"
n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
b[:n] = "eader!"
n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
b[:n] = ""
```

Each ASCII letter in this example is 1 byte, that's why we see it reads 8 chars then 6


## Exercise: Readers
Implement a `Reader` type that emits an infinite stream of the ASCII character `'A'`

```Go
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r *MyReader) Read(b []byte) (n int, err error) {
	//Iterate over the lenght of bytes, assign all to A
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
 }


func main() {
	//Can get the pointer here since it's a struct
	reader.Validate(&MyReader{})
}
```

## Exercise: rot13Reader

A common pattern is an [io.Reader](https://go.dev/pkg/io/#Reader) that wraps another `io.Reader`, modifying the stream in some way.

For example, the [gzip.NewReader](https://go.dev/pkg/compress/gzip/#NewReader) function takes an `io.Reader` (a stream of compressed data) and returns a `*gzip.Reader` that also implements `io.Reader` (a stream of the decompressed data).

Implement a `rot13Reader` that implements `io.Reader` and reads from an `io.Reader`, modifying the stream by applying the [rot13](https://en.wikipedia.org/wiki/ROT13) substitution cipher to all alphabetical characters.

The `rot13Reader` type is provided for you. Make it an `io.Reader` by implementing its `Read` method.

```Go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

//ROT13 is a simple letter substitution cipher that replaces a letter with the //13th letter after it in the Latin alphabet.
func rot13(c byte) byte {
	switch {
	case c >= 'A' && c <= 'Z':
		return 'A' + (c-'A'+13)%26
	case c >= 'a' && c <= 'z':
		return 'a' + (c-'a'+13)%26
	default:
		return c
	}
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

```