package main

import (
	"net/http"
	"sync"
	"time"
)

// scheduleWork receives channel with input urls and
// schedules them to goroutines wrapping
func scheduleWork(urls chan string, output chan *Result,
	concurrency int, timeout time.Duration) {
	defer close(output)
	tokens := make(chan interface{}, concurrency)

	for i := 0; i < concurrency; i++ {
		tokens <- true
	}

	client := &http.Client{
		Timeout: timeout,
	}
	var wg sync.WaitGroup
	for url := range urls {
		// waiting for being allowed
		// to create new goroutines
		<-tokens
		wg.Add(1)
		go func(url string) {
			defer func() { tokens <- true; wg.Done() }()
			output <- process(url, client)
		}(url)
	}
	wg.Wait()
}

// CheckURL downloads URL, counts amount of "go" line
// and returns *Result
func process(url string, client *http.Client) *Result {
	result := &Result{
		URL: url,
	}

	response, err := client.Get(url)
	if err != nil {
		result.Error = err
		return result
	}

	// Maybe we need to use closure here
	// and check if body was closed correctly
	// but for ReadCloser it doesn't make much sense
	defer response.Body.Close()

	// we can download all body to the string here
	// and use strings.Count,
	// but if it's big enough we are risking run out of memory
	// so it's better to process it as reader
	count, err := CountGos(response.Body)
	if err != nil {
		result.Error = err
		return result
	}

	result.Count = count
	return result
}
