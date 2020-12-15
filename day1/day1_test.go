package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	filePath := "./testInput.txt"
	pair := solveDay1(filePath)
	want := Pair{299, 1721, 514579}
	if want != pair {
		t.Fatalf(`Want = %v, expected %v`, want, pair)
	}
}
