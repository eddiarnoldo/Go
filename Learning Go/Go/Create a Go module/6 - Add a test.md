In Go you can use the `testing` package to create your tests. The convention is to name your tests with the following suffix `_test.go`

This way you can simply run `go test` or `go test -v` for verbose mode.

Here is a sample of the output of the verbose mode
```bash
eddiarnoldo@Eddy:~/programming/Go/modules/greetings$ go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      example.com/greetings   0.002s
```


## 1st Test Explanation

Here is the code of our tests created for the new module
```go
func TestHelloName(t *testing.T) {
	name := "Leonel"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Leonel")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Leonel") = %q, %q, want match for %#q, nil`, msg, err, want)
	}
}
```

- `t` is a pointer to [testing.T type](https://go.dev/pkg/testing/#T) which exposes some functions that can be used for your tests:
	- Fail
	- FailNow
	- Log | Logf
- `regexp.MustCompile(`\b` + name + `\b`)
	- This portion is creating a new `*Regex` type which is basically a regex object we can use to match the string
	- \b this is a word boundary (matches the word without it being part of other word)
- Lastly we use the `t` pointer to log an error in the event of a fail to match or if the func returns an error

> Strings in Go can be compared with `==`

## Test 2 explanation

```go
// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
```

This one is more simple here we are just confirming that our Module function returns an empty string if we pass and empty name to receive a greeting.

