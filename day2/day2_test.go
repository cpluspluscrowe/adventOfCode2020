package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	filePath := "./testInput.txt"
	validCount := getDay2ValidCounts(filePath)
	want := 2
	if want != validCount {
		t.Fatalf(`Want = %v, got %v`, want, validCount)
	}
}
