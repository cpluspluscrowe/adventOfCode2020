package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func readFile(filePath string) string {
	contents, _ := ioutil.ReadFile(filePath)
	var s string = string(contents)
	return s
}

type Data struct {
	low, high    int
	letter, text string
}

func convertToNumber(toConvert string) (int, error) {
	number, err := strconv.Atoi(toConvert)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func parseLine(line string) (Data, error) {
	split1 := strings.Split(line, "-")
	low, error1 := convertToNumber(split1[0])
	if error1 != nil {
		return Data{}, error1
	}		
	split2 := strings.Split(split1[1], " ")
	high, error2 := convertToNumber(split2[0])
	if error2 != nil {
		return Data{}, error2
	}			
	letter := strings.Replace(split2[1], ":", "", 1)
	text := split2[2]
	data := Data{low, high, letter, text}
	return data, nil
}

func parseLines(lines []string) []Data {
	result := []Data{}
	for _, line := range lines {
		parsed, err := parseLine(line)
		if err == nil{
			result = append(result, parsed)
		}
	}
	return result
}

func isValidPart1(data Data) bool {
	count := 0
	for _, letter := range data.text {
		if string(letter) == data.letter {
			count += 1
		}
	}
	if count < data.low || count > data.high {
		return false
	}
	return true
}

func isValid(data Data) bool {
	position1 := data.low - 1
	position2 := data.high - 1
	letter := data.letter
	text := data.text
	t1 := string(text[position1]) == letter
	t2 := string(text[position2]) == letter
	if (t1 && !t2) || (t2 && !t1) {
		return true
	}else{
		return false
	}
}

func countValids(lines [] Data) int {
	countOfValidCases := 0
	for _, line := range lines {
		valid := isValid(line)
		if valid {
			countOfValidCases += 1
		}
	}
	return countOfValidCases
}

func getDay2ValidCounts(filePath string) int {
	text := readFile(filePath)
	lines := strings.Split(text, "\n")
	data := parseLines(lines)
	validCount := countValids(data)
	return validCount
}

func main() {
	filePath := "./input.txt"
	validCount := getDay2ValidCounts(filePath)
	fmt.Println(validCount)
}
