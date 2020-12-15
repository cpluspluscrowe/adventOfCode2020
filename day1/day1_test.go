package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	filePath := "./testInput.txt"
	pair := solveDay1(filePath)
	want := Pair{366, 675, 979, 241861950}
	if want != pair {
		t.Fatalf(`Want = %v, got %v`, want, pair)
	}
}
