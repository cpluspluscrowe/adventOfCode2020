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

func main() {
	filePath := "./input.txt"
	text := readFile(filePath)
	lines := strings.Split(text, "\n")
	data := parseLines(lines)
	fmt.Println(data)
}
