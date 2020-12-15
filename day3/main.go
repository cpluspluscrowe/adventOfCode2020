package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(filePath string) string {
	contents, _ := ioutil.ReadFile(filePath)
	var s string = string(contents)
	return s
}

func createTreeMap(line string) map[int]bool {
	result := make(map[int]bool)
	for i, character := range line {
		if string(character) == "#" {
			result[i] = true
		}
	}
	return result
}

func joinTreeMappings(lines []string) map[int]map[int]bool {
	joinedTreeMappings := make(map[int]map[int]bool)
	for row, line := range lines {
		treeMapping := createTreeMap(line)
		if _, ok := joinedTreeMappings[row]; !ok {
			joinedTreeMappings[row] = make(map[int]bool)
		}
		for treeIndex, _ := range treeMapping {
			joinedTreeMappings[row][treeIndex] = true
		}
	}
	return joinedTreeMappings
}

type TobogganPath struct {
	right, down int
}

type Position struct {
	row, column int
}

func createTobogganPath(tp TobogganPath, depth int) []Position {
	willVisit := []Position{}
	rightShift := 0
	downShift := 0
	for downShift < depth {
		rightShift += tp.right
		downShift += tp.down
		if downShift >= depth {
			continue
		}
		position := Position{downShift, rightShift}
		willVisit = append(willVisit, position)

	}
	return willVisit
}

func countTreeHits(tpPath []Position, trees map[int]map[int]bool, width int) int {
	hitTreeCount := 0
	for _, position := range tpPath {
		column := position.column % width
		if _, ok := trees[position.row][column]; ok {
			hitTreeCount += 1
		}
	}
	return hitTreeCount
}

func multipleTreeHitsForSlopes(slopes []TobogganPath, trees map[int]map[int]bool, depth int, width int) int {
	multipliedTreeHits := 1
	for _, tp := range slopes {
		tpPath := createTobogganPath(tp, depth)
		hitTreeCount := countTreeHits(tpPath, trees, width)
		fmt.Println(hitTreeCount)
		multipliedTreeHits *= hitTreeCount
	}
	return multipliedTreeHits
}

func getDay3TreeHitCount(filePath string) int {
	text := readFile(filePath)
	lines := strings.Split(text, "\n")
	width := len(lines[0])
	trees := joinTreeMappings(lines)
	depth := len(trees)
	fmt.Println(depth, width)
	slopesToCheck := []TobogganPath{TobogganPath{1, 1},TobogganPath{3,  1},TobogganPath{5,  1},TobogganPath{7,  1},TobogganPath{1,  2}}
	multipliedTreeHits := multipleTreeHitsForSlopes(slopesToCheck, trees, depth, width)
	return multipliedTreeHits
}

func main() {
	filePath := "./input.txt"
	multipliedTreeHits := getDay3TreeHitCount(filePath)
	fmt.Println(multipliedTreeHits)
}
