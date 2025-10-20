

Switchs in Go act more like if statements, they break by default and they are more flexible on what you can switch case on 

```Go
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```


>"Another important difference is that Go's switch is more flexible: cases can be variables or expressions (not just constants), and you can switch on any comparable type (not just integers)."

can switch on variables/expressions
```Go
	// Go ✅ - Java ❌
	x := 10
	y := 20
	
	switch value {
	case x:           // Variable as case
	case y + 5:       // Expression as case  
	case len("hello"): // Function call as case
	}
```

Can switch on strings
```Go
	// Go ✅ - Java ❌ (well, Java added this in Java 7, but traditionally no)
	switch name {
	case "alice":
	    fmt.Println("Hello Alice")
	case "bob":
	    fmt.Println("Hello Bob")
	}
```

Multiple values
```Go
	// Go ✅ - Java ❌
	switch day {
	case "saturday", "sunday":
	    fmt.Println("Weekend!")
	case "monday", "tuesday", "wednesday", "thursday", "friday":
	    fmt.Println("Weekday")
	}
```

```Go
	// Go ✅ - Java ❌
	type Status bool
	var isReady Status = true
	
	switch isReady {
	case true:
	    fmt.Println("Ready")
	case false:
	    fmt.Println("Not ready")
	}
```


```Go
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

> ^ this is interesting since we use a constant, and we use the case as if's to show a message if Saturday is still away from us


## Switch with no condition = switch true {}

>This construct can be a clean way to write long if-then-else chains.

```Go
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```
