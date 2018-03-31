package main

import (
	"bufio"
	"io"
)

// CountGos counts amount of substring "go" in input io.Reader.
// It returns this amount and error in caseif something bad happened.
// "go" is hardcoded, but for searching substrings with length > 2
// we would need to change algorithm, and it's not required by the task.
// This realization has an issue, if input is big and does not contain 'o'
// we would risk running out of memory anyway.
func CountGos(input io.Reader) (int, error) {
	count := 0
	reader := bufio.NewReader(input)

	for {
		s, err := reader.ReadString('o')
		if err == io.EOF {
			break
		} else if err != nil {
			return count, err
		}

		l := len(s)
		if l < 2 {
			continue
		}

		if s[l-2] == 'G' {
			count++
		}
	}

	return count, nil
}
