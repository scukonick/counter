package main

// Result represents result of
// check of one URL
type Result struct {
	URL   string
	Count int
	Error error
}
