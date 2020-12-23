package main

import (
	"testing"
)

/*func TestRowIds(t *testing.T) {
	filePath := "./inputTest.txt"
	sum := getDaySixPartOneSummation(filePath)
	want := 11
	if want != sum {
		t.Fatalf(`Want = %v, got %v`, want, sum)
	}
        }*/

func TestRowIds(t *testing.T) {
	filePath := "./inputTest.txt"
	sum := getDaySixPartOneSummation(filePath)
	want := 6
	if want != sum {
		t.Fatalf(`Want = %v, got %v`, want, sum)
	}
}
