## Tiny main abstraction
Do not just throw errors from main directly call a function from main

```go
func main(){
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	} 
}

func run() error {
	db, dbtidy, err := setupDatabase()
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer dbtidy()
	srv := &server{
		db: db,
	}
	// more stuff
}
```

## The server struct
- Represent the component as a struct
- Shared dependencies as fields
- No global state

```go
type server struct {
	db *someDatabase
	router *someRouter
	email EmailSender
}
```


## 5 words of advice for new/intermediate Go devs
 > never use global variables
 
 
## Constructor for server
```Go
func newServer() *server {
	s := &server{}
	s.routes()
	return s
}
``` 

- I ofter end with a constructor
- **Don't setup dependencies here**, remember you might want to use this in test code
- You can take dependencies as arguments if there are not many

## If you're doing you http server make it an http.Handler

`http.Handler` is an interface and it can be simply used by implementing the method

```Go
func (s *server) ServeHTTP(w http.ResponseWritter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
```

- Implement `ServeHTTP` in your server so you can turn your server into an `http.Handler`
- Use your **server** anywhere where you can use an `http.Handler`
- Just pass execution to your router (don't hide logic in here, use middleware instead as it's more clear) 
	- This refers to the line `s.router.ServeHTTP(w, r)`

## routes.go
```Go
package main

func (s *server) routes() {
	s.router.Get("/api/", s.handleAPI())
	s.router.Get("/about" s.handleAbout())
	s.router.Get("/", s.handleIndex())
}
```

- One place for all routes
- Most code maintenance starts with a URL so it's handy to have one place to look

## Handlers that hang off the server

```Go
func (s *server) handleSomething() http.HandlerFunc {
	//put some code here
}
```

- Handlers are methods on the server, which give them access to dependencies via `s`
- Remember , other handlers have access to `s` too, be careful with data races.


## Naming handler methods

```Go
handleTasksCreate
handleTasksUpdate
handleTasksGet

handleAuthLogin
handleAuthLogout
```

This naming will help you since it uses alphabetical order to keep your methods sorted and kind of groups if you use a documentation generator (group related to functionalty)

## Return the handler
```Go
func (s *server) handleSomething() http.HandlerFunc {
	something := prepareSomething()
	return func(w http.ResponseWritter, r *http.Request) {
		//use "something"
	}
}
```

- Allows handler specific setup

## Take arguments for handler-specific dependencies
if you have any specific dependencies for a couple of handlers that you don't want to put in your server type you can pass them in your handler return function

```Go
func (s *server) handleGreeting(format string) http.HandlerFunc {
	return func(w http.ResponseWritter, r *http.Request) {
		fmt.FprintF(w, format, r.FormValue("name"))
	}
}

s.router.HandleFunc("/one", s.handleGreeting("Hola %s"))
s.router.HandleFunc("/two", s.handleGreeting("Hello %s"))
```

### More examples of this

```Go
handleTemplate(template *template.Template) http.HandlerFunc

handleRandomQuote(q Quoter, r *rand.Rand) http.HandlerFunc

handleSendMagicLink(e EmailSender) http.HandlerFunc
```

- This makes it easy to know the dependencies needed by every handler to do its job
- Type safety and compile time of the app helps to make sure you provide the needed arguments


## Too Big? Have many servers

```Go

//people.go
type serverPeople struct {
	db *mydatabase
	emailSender EmailSender
}

//comments.go
type serverComments struct {
	db *mydatabase
}
```

## HandlerFunc over Handler
```Go
func (s *server) handleSomething() http.HandlerFunc {
	return func(w http.ResponseWritter, r *http.Request) {
	}
}
```

- `http.HandlerFunc` is a custom type that implements the interface `http.Handler` so they are more or less interchangeable
- Pick whichever is easier to read in your case
- Occasionally you might deal with converting between them

## Middleware are just Go functions
```Go
func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc () {
		return func (w http.ResponseWritter, r *http.Request) {
			if !currentUser(r).isAdmin {
				http.notFound(w, r)
				return
			}
			h(w, r)
		}
}
```

^ This is cool since we can take a handlerFunc which we will call the `"original"` func and run some code before or after it, in this example we take `h` and check if it's an admin and serve if it's or return a not found http response if not an admin.

**Notice this is also a method on the server!**

- Take an `http.HandlerFunc` and return a new one.
- Run code before/after the wrapped handler
- Or choose not to call the wrapped handler at all

> You can use this same pattern to do authentication, logging 

## Wire Middleware up in `routes.go`

```Go
package main

func (s *server) routes() {
	s.router.Get("/api/", s.handleAPI())
	s.router.Get("/about", s.handleAbout())
	s.router.Get("/", s.handleIndex())
	s.router.Get("/admin", s.adminOnly(s.handleAdminIndex()))
	// Notice how we use the pattern above to protect our admin route
}
```

> `routes.go` becomes a high level map of the service

# Dealing with data
## Respond helper

```Go
func (s *server) respond(w http.ResponseWritter, r *http.Response, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		//TODO handle error
	}
		
}
```

> As of Go 1.18 you can use `data any` 


Don't over abstract stuff start simple, e.g don't create a function to prevent repetitive code until you really needed e.g here we have a respond helper method on the server which can evolve

- Abstract responding and do the barebones initially
- Later you can make this more sophisticated (if needed)

## Decoding helper
