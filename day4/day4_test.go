package main

import (
	"testing"
)

func TestValidPassportsCount(t *testing.T) {
	filePath := "./inputTest.txt"
	countOfValidPassports := getDay4ValidPassportCount(filePath)
	want := 2
	if want != countOfValidPassports {
		t.Fatalf(`Want = %v, got %v`, want, countOfValidPassports)
	}
}

func TestValidityOfPassportStrings(t *testing.T) {
	if validBirthYear("2002") != true {
		t.Fatalf(`Want %v`, true)
	}
	if validBirthYear("2003") != false {
		t.Fatalf(`Want %v`, false)
	}
	if validHeight("60in") != true {
		t.Fatalf(`Want %v`, true)
	}
	if validHeight("190cm") != true {
		t.Fatalf(`Want %v`, true)
	}
	if validHeight("190in") != false {
		t.Fatalf(`Want %v`, false)
	}
	if validHeight("190") != false {
		t.Fatalf(`Want %v`, false)
	}

	if validHairColor("#123abc") != true {
		t.Fatalf(`Want %v`, true)
	}
	if validHairColor("#123abz") != false {
		t.Fatalf(`Failed`)
	}
	if validHairColor("123abc") != false {
		t.Fatalf(`Failed`)
	}
	if validPassportId("000000001") != true {
		t.Fatalf(`Failed`)
	}
	if validPassportId("0123456789") != false {
		t.Fatalf(`Failed`)
	}
	if validEyeColor("brn") != true {
		t.Fatalf(`Failed`)
	}
	if validEyeColor("wat") != false {
		t.Fatalf(`Failed`)
	}
}
