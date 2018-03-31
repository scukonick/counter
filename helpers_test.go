package main

import (
	"strings"
	"testing"
)

func TestCountGos(t *testing.T) {
	s := "Go1231231223GoaaaaaasdasdGoGogG__ooooo"

	r := strings.NewReader(s)

	c, err := CountGos(r)
	if err != nil {
		t.Errorf("Error should be nil if everything is ok: %v", err)
	}
	if c != 4 {
		t.Error("CountGos should count go correctly")
	}

	s = "Go12инарусскомнемногоGogg__ooooo"
	r = strings.NewReader(s)

	c, err = CountGos(r)
	if err != nil {
		t.Errorf("Error should be nil if everything is ok: %v", err)
	}
	if c != 2 {
		t.Error("CountGos should count go correctly")
	}

	s = "o"
	r = strings.NewReader(s)

	c, err = CountGos(r)
	if err != nil {
		t.Errorf("Error should be nil if everything is ok: %v", err)
	}
	if c != 0 {
		t.Error("CountGos should count go correctly")
	}

	s = ""
	r = strings.NewReader(s)

	c, err = CountGos(r)
	if err != nil {
		t.Errorf("Error should be nil if everything is ok: %v", err)
	}
	if c != 0 {
		t.Error("CountGos should count go correctly")
	}
}
