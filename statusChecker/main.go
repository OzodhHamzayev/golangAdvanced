package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL string
	Status int
	Duration time.Duration
	Err error
}

func check(url string) Result {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)
	if err != nil {
		return Result{URL: url, Duration: duration, Err: err}
	}
	resp.Body.Close()
	return Result{URL: url, Status: resp.StatusCode, Duration: duration}
}	
func worker(id int, jobs <- chan string, results chan <- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		results <- check(url)
	}
}


func main( ) {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://go.dev",
		"https://golang.org",
		"https://example.com",
	}
	jobs := make(chan string, len(urls))
	results := make(chan Result)

	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for _, url := range urls {
		jobs <- url
	} 
	close(jobs)


	go func ()  {
		wg.Wait()
		close(results)
	}()


	for r := range results {
		if r.Err != nil {
			fmt.Printf("%-25s XATO: %v\n", r.URL, r.Err)
			continue
		}
		fmt.Printf("%-25s %d  %v\n", r.URL, r.Status, r.Duration)
	}



	fmt.Println("Full time:", time.Since(start))
}



