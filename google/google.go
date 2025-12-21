package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q", kind, query))
	}
}

var (
	Web = fakeSearch("web")
	Img = fakeSearch("image")
	Vid = fakeSearch("video")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("gundam wing")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Printf("Search1 took %s\n", elapsed)
	fmt.Println("====== \n")

	//Here now we can see it only waits for the slowests search to complete
	start2 := time.Now()
	results2 := Google2("samurai x")
	elapsed2 := time.Since(start2)

	fmt.Println(results2)
	fmt.Printf("Search2 took %s\n", elapsed2)
	fmt.Println("====== \n")
	//This one has a slight improvement it does not wait for those searches that take longer than 80 ms
	start3 := time.Now()
	results3 := Google21("One Piece x")
	elapsed3 := time.Since(start3)

	fmt.Println(results3)
	fmt.Printf("Search2 took %s\n", elapsed3)

}

// This one makes sequential calls so each search has to wait for the previous one to complete
func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Img(query))
	results = append(results, Vid(query))
	return
}

// This one makes concurrent calls so all searches happen simultaneously, we push a value or send a value to the channel when each search completes
func Google2(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Img(query) }()
	go func() { c <- Vid(query) }()

	for i := 0; i < 3; i++ {
		results = append(results, <-c)
	}
	return
}

// This one has a slight improvement it does not wait for those searches that take longer than 80 ms
func Google21(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Img(query) }()
	go func() { c <- Vid(query) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case res := <-c:
			results = append(results, res)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	//this is a no params return, this can be done because results is named in the function signature
	return
}
