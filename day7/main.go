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
	number := strings.Split(line, " ")[0]
	count := convertToNumber(number)
	remainder := strings.Trim(strings.Replace(line, number, "", 1), " ")
	bagReference := BagReference{remainder, count}
	return bagReference

}

func parseLine(line string) OuterBag {
	bagsToBag := strings.Replace(line, "bags", "bag", -1)
	zeroTo0 := strings.Replace(strings.Replace(bagsToBag, "no", "0", 1), ".", "", 1)
	toSplitOn := " contain "
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

func createBagMap(bagContents []OuterBag) map[string]OuterBag {
	bagMap := make(map[string]OuterBag)
	for _, outerBag := range bagContents {
		color := outerBag.color
		bagMap[color] = outerBag
	}
	return bagMap
}

func canContainYellowBags(bagColor string, bagMap map[string]OuterBag) bool {
	if bagColor == "shiny gold bag" {
		return true
	}
	outerBag := bagMap[bagColor] // throw error if this does not exist
	goldBagCount := 0
	for _, innerBag := range outerBag.contains {
		if canContainYellowBags(innerBag.color, bagMap) { // * innerBag.count
			goldBagCount += 1
		}
	}
	return goldBagCount > 0
}

func searchOuterBags(outerBags []OuterBag, bagMap map[string]OuterBag) int {
	goldBagCount := 0
	for _, outerBag := range outerBags {
		if outerBag.color != "shiny gold bag" {
			if canContainYellowBags(outerBag.color, bagMap) {
				goldBagCount += 1
			}
		}
	}
	return goldBagCount
}

func main() {
	var filePath string = "./input.txt"
	var text string = readFile(filePath)
	var lines []string = strings.Split(text, "\n")
	var removeLastLine []string = lines[:len(lines)-1]
	outerBags := parseLines(removeLastLine)
	bagMap := createBagMap(outerBags)
	goldBagCount := searchOuterBags(outerBags, bagMap)
	fmt.Println(goldBagCount)
}
