package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	filePath := "./inputTest.txt"
	validCount := getDay3TreeHitCount(filePath)
	want := 7
	if want != validCount {
		t.Fatalf(`Want = %v, got %v`, want, validCount)
	}
}
