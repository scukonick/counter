package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	concurrency := flag.Int("c", 5, "Maximum number of concurrent goroutines")
	timeoutH := flag.String("t", "5s", "Timeout for requests")
	flag.Parse()

	timeout, err := time.ParseDuration(*timeoutH)
	if err != nil {
		log.Fatalf("Failed to parse timeout: %v", err)
	}
	if *concurrency <= 0 {
		log.Fatalf("Concurrency should be positive integer: %d", *concurrency)
	}

	totalCount := 0
	// capacity may be used for buffering output
	results := make(chan *Result, 10)
	urls := make(chan string)

	go readInput(os.Stdin, urls)
	go scheduleWork(urls, results, *concurrency, timeout)

	for result := range results {
		if result.Error != nil {
			fmt.Printf("URL: %s processed with error: %v\n", result.URL, result.Error)
		} else {
			// Actually, we can print this inside goroutines
			// using lgg module which is thread safe, so we would be able to pass
			// to result channel only
			// counts instead of Result structures,
			// but from the example of output in task it looked like
			// it was done by simple fmt.Printf
			fmt.Printf("Count for %s: %d\n", result.URL, result.Count)
			totalCount += result.Count
		}
	}

	fmt.Printf("Total count: %d\n", totalCount)
}
