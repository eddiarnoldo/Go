Now that we have our module tested, we need to run it of since we cannot keep using `go run .` forever.

- `go build`
	- This will create an executable which we can run like ./hello
- `go install`
	- This requires more steps but will allow us to run the executable from anywhere
		- For this we need to retrieve the path ` go list -f '{{.Target}}'` then we need to add that into our $PATH
		- After that we can just run hello

```bash
eddiarnoldo@Eddy:~/programming/Go/modules/hello$ hello
Great to see you, Eddy!
map[Eddy:Hail, Eddy well met! Leo:Hail, Leo well met! Rocio:Hi, Rocio. Welcome]
```

