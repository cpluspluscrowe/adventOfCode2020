package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(filePath string) string {
	contents, _ := ioutil.ReadFile(filePath)
	var text string = string(contents)
	return text
}

// light red bags contain 1 bright white bag, 2 muted yellow bags.

func convertToNumber(toConvert string) int {
	number, _ := strconv.Atoi(toConvert)
	return number
}

type BagReference struct {
	color string
	count int
}

type OuterBag struct {
	color    string
	contains []BagReference
}

func parseInnerBag(line string) BagReference {
	//3 faded blue bags
	number := strings.Split(line, " ")[0]
	count := convertToNumber(number)
	remainder := strings.Trim(strings.Replace(line, number, "", 1), " ")
	bagReference := BagReference{remainder, count}
	return bagReference

}

func parseLine(line string) OuterBag {
	zeroTo0 := strings.Replace(strings.Replace(line, "no", "0", 1), ".", "", 1)
	toSplitOn := " bags contain "
	bagColorAndDetails := strings.Split(zeroTo0, toSplitOn)
	bag := OuterBag{}
	bag.color = bagColorAndDetails[0]
	bag.contains = []BagReference{}
	byComma := strings.Split(bagColorAndDetails[1], ", ")
	for _, innerBag := range byComma {
		bagReference := parseInnerBag(innerBag)
		bag.contains = append(bag.contains, bagReference)
	}
	return bag
}

func parseLines(lines []string) []OuterBag {
	result := []OuterBag{}
	for _, line := range lines {
		bag := parseLine(line)
		result = append(result, bag)
	}
	return result
}

func main() {
	var filePath string = "./input.txt"
	var text string = readFile(filePath)
	var lines []string = strings.Split(text, "\n")
	var removeLastLine []string = lines[:len(lines)-1]
	outerBags := parseLines(removeLastLine)
	fmt.Println(outerBags)
}
