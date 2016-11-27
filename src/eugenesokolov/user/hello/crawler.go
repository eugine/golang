package main

import "fmt"

type FetchResult struct {
	url string 
	body string
	urls []string
	err  error
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (result FetchResult)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) map[string]FetchResult {
	fetchedResults := make(map[string]FetchResult)
	urls := []string{url}
	for i := depth; i > 0; i-- {
		onlyNewUrls := []string{}
		for _, u := range urls {
			if _, ok := fetchedResults[u]; !ok {
				onlyNewUrls = append(onlyNewUrls, u)
			}
		}
		chans := []chan FetchResult{}
		for _, u := range onlyNewUrls {
			ch := make(chan FetchResult)
			go fetch(u, fetcher, ch)
			chans = append(chans, ch)
		}

		urls = []string{}
		for _, ch := range chans {
			result := <- ch
			fetchedResults[result.url] = result
			for _, u := range result.urls {
				urls = append(urls, u)
			}
		}
		if (len(urls) == 0) {
			break
		}
	}
	return fetchedResults
}

func fetch(url string, fetcher Fetcher, ch chan FetchResult) {
	fmt.Println("-- quering: " + url)
	ch <- fetcher.Fetch(url)
}

func main() {
	result := Crawl("http://golang.org/", 4, fetcher)
	fmt.Println("Result:")
	for url, res := range result {
		if (res.err == nil) {
			fmt.Printf("%v == %v\n", url, res.body)
		} else {
			fmt.Printf("ERROR: %v\n", res.err)
		}
		
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) FetchResult {
	if res, ok := f[url]; ok {
		return FetchResult{url, res.body, res.urls, nil}
	}
	return FetchResult{url, "", nil, fmt.Errorf("not found: %s", url)}
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
