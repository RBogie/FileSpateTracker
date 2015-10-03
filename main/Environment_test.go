package main

import "testing"

func TestAverage(t *testing.T) {
	i := 40
	j := 2
	answer := i + j
	if answer != 42 {
		t.Error("Expected 42 as the answer to everything, but got ", answer)
	}
}
