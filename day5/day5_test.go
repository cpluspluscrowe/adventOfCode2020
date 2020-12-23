package main

import (
	"testing"
)

func TestRowIds(t *testing.T) {
	filePath := "./inputTest.txt"
	maxRowId := getMaximumRowId(filePath)
	want := 820
	if want != maxRowId {
		t.Fatalf(`Want = %v, got %v`, want, maxRowId)
	}
}
