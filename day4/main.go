package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filePath string) string {
	contents, _ := ioutil.ReadFile(filePath)
	var s string = string(contents)
	return s
}

func splitByDelimiters(text string) [][]string {
	result := [][]string{}
	chunks := strings.Split(text, "\n\n")
	for _, chunk := range chunks {
		pieces := []string{}
		split1 := strings.Split(chunk, "\n")
		for _, piece1 := range split1 {
			split2 := strings.Split(piece1, " ")
			for _, piece2 := range split2 {
				if piece2 != "" {
					pieces = append(pieces, piece2)
				}
			}
		}
		result = append(result, pieces)
	}
	return result
}

func convertPassportStringToMap(passport []string) map[string]string {
	passportMappings := make(map[string]string)
	for _, piece := range passport {
		keyValue := strings.Split(piece, ":")
		if len(keyValue) != 2 {
			panic("piece does not contain colon, invalid piece!")
		}
		passportMappings[keyValue[0]] = keyValue[1]
	}
	return passportMappings
}

func processPassportStrings(passports [][]string) []map[string]string {
	passportMaps := []map[string]string{}
	for _, passport := range passports {
		passportMap := convertPassportStringToMap(passport)
		passportMaps = append(passportMaps, passportMap)
	}
	return passportMaps
}

func isValidPassportOld(passport map[string]string) bool {
	necessaryKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, necessaryKey := range necessaryKeys {
		if _, ok := passport[necessaryKey]; !ok {
			return false
		}
	}
	return true
}

func convertToNumber(toConvert string) (int, error) {
	number, err := strconv.Atoi(toConvert)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func validBirthYear(birthYearValue string) bool {
	number, err := convertToNumber(birthYearValue)
	if err != nil {
		return false
	}
	if number >= 1920 && number <= 2002 {
		return true
	}
	return false
}

func validIssueYear(issueYear string) bool {
	number, err := convertToNumber(issueYear)
	if err != nil {
		return false
	}
	if number >= 2010 && number <= 2020 {
		return true
	}
	return false
}

func validExpirationYear(expirationYear string) bool {
	number, err := convertToNumber(expirationYear)
	if err != nil {
		return false
	}
	if number >= 2020 && number <= 2030 {
		return true
	}
	return false
}

func validHeight(height string) bool {
	validInches := strings.Contains(height, "in")
	validCentimeters := strings.Contains(height, "cm")
	valid := validInches || validCentimeters
	if !valid {
		return false
	}
	leftOver := strings.Replace(strings.Replace(height, "in", "", 1), "cm", "", 1)
	number, err := convertToNumber(leftOver)
	if err != nil {
		return false
	}
	if number >= 59 && number <= 76 && validInches {
		return true
	}
	if number >= 150 && number <= 193 && validCentimeters {
		return true
	}
	return false
}

func validHairColor(hairColor string) bool {
	if len(hairColor) == 0 {
		return false
	}
	valid := string(hairColor[0]) == "#"
	if !valid {
		return false
	}
	remainder := hairColor[1:len(hairColor)]
	pattern := "([0-9]|[a-f]){6}"
	matched, err := regexp.MatchString(pattern, remainder)
	if err != nil {
		return false
	}
	return matched
}

func validEyeColor(eyeColor string) bool {
	allowableCount := 0
	allowableColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, color := range allowableColors {
		if eyeColor == color {
			allowableCount += 1
		}
	}
	valid := allowableCount == 1
	return valid
}

func validPassportId(passportId string) bool {
	pattern := "[0-9]+"
	matched, err := regexp.MatchString(pattern, passportId)
	if err != nil {
		return false
	}
	if len(passportId) == 9 && matched {
		return true
	}
	return false
}

func isValidPassport(passport map[string]string) bool {
	iyr := validIssueYear(passport["iyr"])
	byr := validBirthYear(passport["byr"])
	eyr := validEyeColor(passport["ecl"])
	hgt := validPassportId(passport["pid"])
	hcl := validHairColor(passport["hcl"])
	ecl := validHeight(passport["hgt"])
	pid := validExpirationYear(passport["eyr"])
	if byr && iyr && eyr && hgt && hcl && ecl && pid {
		return true
	}
	return false
}

func countValidPassports(passports []map[string]string) int {
	count := 0
	for _, passport := range passports {
		if isValidPassport(passport) {
			count += 1
		}
	}
	return count
}

func getDay4ValidPassportCount(filePath string) int {
	text := readFile(filePath)
	passportText := splitByDelimiters(text)
	passportMaps := processPassportStrings(passportText)
	countOfValidPassports := countValidPassports(passportMaps)
	return countOfValidPassports
}

func main() {
	filePath := "./input.txt"
	countOfValidPassports := getDay4ValidPassportCount(filePath)
	fmt.Println(countOfValidPassports)
}
