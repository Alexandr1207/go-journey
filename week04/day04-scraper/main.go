package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Result struct {
	URL      string
	Status   int
	Size     int
	Title    string
	Duration time.Duration
	Err      error
}

var titleRe = regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)

func fetch(client *http.Client, url string) Result {
	start := time.Now()

	res := Result{URL: url}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		res.Duration = time.Since(start)
		return res
	}
	defer resp.Body.Close()
	res.Status = resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		res.Duration = time.Since(start)
		return res
	}

	if m := titleRe.FindSubmatch(body); m != nil {
		res.Title = strings.TrimSpace(string(m[1]))
	}
	res.Size = len(body)
	res.Duration = time.Since(start)

	return res
}

func printResults(results []Result) {
	for _, r := range results {
		if r.Err != nil {
			fmt.Printf("%-35s ERROR  %-8v %v\n", r.URL, r.Duration.Round(time.Millisecond), r.Err)
			continue
		}
		fmt.Printf("%-35s %d  %6d B  %-8v %s\n", r.URL, r.Status, r.Size, r.Duration.Round(time.Millisecond), r.Title)
	}
}

func runSequential(client *http.Client, urls []string) {
	results := make([]Result, len(urls))
	for i, url := range urls {
		results[i] = fetch(client, url)
	}

	printResults(results)
}

func runConcurrent(client *http.Client, urls []string) {
	var wg sync.WaitGroup
	results := make([]Result, len(urls))
	for i, url := range urls {
		wg.Go(func() {
			results[i] = fetch(client, url) // no races
		})

	}
	wg.Wait()

	printResults(results)
}

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	urls := []string{"https://books.toscrape.com", "https://quotes.toscrape.com", "https://example.com", "https://go.dev", "https://gobyexample.com", "https://httpstat.us/404", "https://localhost:1", "https://httpbin.org/delay/10"}

	start := time.Now()
	runSequential(client, urls)
	end := time.Since(start)
	fmt.Println("Sequential completed with time: ", end) // 8.4217687s

	start = time.Now()
	runConcurrent(client, urls)
	end = time.Since(start)
	fmt.Println("Concurrent completed with time: ", end) // 5.0018447s

	// Output example:
	// https://books.toscrape.com          200   51294 B  1.363s   All products | Books to Scrape - Sandbox
	// https://quotes.toscrape.com         200   11064 B  1.157s   Quotes to Scrape
	// https://example.com                 200     577 B  309ms    Example Domain
	// https://go.dev                      200   64185 B  407ms    The Go Programming Language
	// https://gobyexample.com             200    6884 B  267ms    Go by Example
	// https://httpstat.us/404             ERROR  729ms    Get "https://httpstat.us/404": EOF
	// https://localhost:1                 ERROR  10ms     Get "https://localhost:1": dial tcp [::1]:1: connectex: No connection could be made because the target machine actively refused it.
	// https://httpbin.org/delay/10        ERROR  5s       Get "https://httpbin.org/delay/10": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
	// Program completed with time:  5.0037172s
}
