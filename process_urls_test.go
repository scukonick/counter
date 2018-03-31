package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestScheduleWork(t *testing.T) {
	goCount := 1
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := ""
		for i := 0; i < goCount; i++ {
			result += "Go"
		}
		goCount++
		fmt.Fprintln(w, result)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	input := make(chan string)
	output := make(chan *Result)
	timeout := time.Second
	concurrency := 10

	go scheduleWork(input, output, concurrency, timeout)

	// testing method
	input <- ts.URL
	result := <-output
	if result.Count != 1 {
		t.Errorf("Go count should be 1 for 1 run: %d", result.Count)
	}

	input <- ts.URL
	result = <-output
	if result.Count != 2 {
		t.Errorf("Go count should be 2 for 2 run: %d", result.Count)
	}

	input <- ts.URL
	result = <-output
	if result.Count != 3 {
		t.Errorf("Go count should be 3 for 3 run: %d", result.Count)
	}

	close(input)
}
