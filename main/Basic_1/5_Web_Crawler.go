package main

import (
	"fmt"
	"sync"
)

type crawledUrls struct {
	mu   sync.Mutex
	urls map[string]bool
}

func (c *crawledUrls) CheckAndAdd(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.urls[url] {
		return false
	}
	c.urls[url] = true
	return true
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, c *crawledUrls, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	if !c.CheckAndAdd(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, c, wg)
	}
	return
}

func main() {
	cache := &crawledUrls{urls: make(map[string]bool)}

	var wg sync.WaitGroup

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, cache, &wg)

	wg.Wait()
	fmt.Println("Crawling finished!")
}

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
	"https://golang.org/cmd/": &fakeResult{
		"Commands",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
