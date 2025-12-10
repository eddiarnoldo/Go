package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	visited map[string]bool
	mu      sync.Mutex
}

func (m *SafeMap) SetVal(key string, val bool) {
	//Lock so only 1 goroutine at a time can access the map
	m.mu.Lock()
	defer m.mu.Unlock()
	m.visited[key] = val
}

func (m *SafeMap) GetVal(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.visited[key]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, status chan bool, safeMap *SafeMap) {

	//if we already visited this url then return and set status chan to true
	if ok := safeMap.GetVal(url); ok {
		status <- true
		return
	}

	safeMap.SetVal(url, true)

	if depth <= 0 {
		status <- false
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		status <- false
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	//Create a slice of channels for each url
	statuses := make([]chan bool, len(urls))

	for index, u := range urls {
		statuses[index] = make(chan bool)
		go Crawl(u, depth-1, fetcher, statuses[index], safeMap)
	}

	//Consume all the child goroutines channels
	for _, childStatus := range statuses {
		<-childStatus
	}

	//This go routine can finish
	status <- true
}

func main() {
	urlMap := SafeMap{visited: make(map[string]bool)}
	status := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, status, &urlMap)
	<-status
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
