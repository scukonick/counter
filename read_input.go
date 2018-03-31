package main

import (
	"bufio"
	"io"
	"log"
)

func readInput(input io.Reader, output chan string) {
	scanner := bufio.NewScanner(input)
	defer close(output)

	for scanner.Scan() {
		url := scanner.Text()
		output <- url
	}

	err := scanner.Err()
	if err != nil {
		log.Fatalf("Unexpected error during reading from input: %v", err)
	}
}
