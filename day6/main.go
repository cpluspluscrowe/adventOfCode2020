package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func readFile(filePath string) string {
	contents, _ := ioutil.ReadFile(filePath)
	var text string = string(contents)
	return text
}

func checkIfLetter(potentialLetter string) bool {
	pattern := "[a-z]"
	matched, _ := regexp.MatchString(pattern, potentialLetter)
	return matched
}

func countLettersInEachGroup(group string) int {
	groupYesCounter := make(map[string]int)
	persons := strings.Split(group, "\n")
	desiredLength := 0
	for _, person := range persons {
		if len(person) > 0 {
			desiredLength += 1

		}
		for _, character := range person {
			as_str := string(character)
			groupYesCounter[as_str] += 1
		}
	}
	count := 0
	for _, occurrenceCount := range groupYesCounter {
		if occurrenceCount == desiredLength {
			count += 1
		}
	}
	return count
}

func countLettersInEachGroupPart1(group string) int {
	groupYesCounter := make(map[string]bool)
	persons := strings.Split(group, "\n")
	for _, person := range persons {
		for _, character := range person {
			as_str := string(character)
			groupYesCounter[as_str] = true
		}
	}
	return len(groupYesCounter)
}

func mapGroupsToYesCounts(groups []string) []int {
	yesCounts := []int{}
	for _, group := range groups {
		yesCount := countLettersInEachGroup(group)
		yesCounts = append(yesCounts, yesCount)
	}
	return yesCounts
}

func sumGroups(yesCounts []int) int {
	sum := 0
	for _, yesCount := range yesCounts {
		sum += yesCount
	}
	return sum
}

func getDaySixPartOneSummation(filePath string) int {
	text := readFile(filePath)
	groups := strings.Split(text, "\n\n")
	yesCounts := mapGroupsToYesCounts(groups)
	sum := sumGroups(yesCounts)
	return sum
}

func main() {
	var filePath string = "./input.txt"
	sum := getDaySixPartOneSummation(filePath)
	fmt.Println(sum)
}
