package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 127 rows, 8 columns

func readFile(filePath string) string {
	contents, _ := ioutil.ReadFile(filePath)
	var s string = string(contents)
	return s
}

// B = upper
func getRow(line string) int {
	rowCharacters := line[:7]
	rowMax := 127
	rowMin := 0
	for _, character := range rowCharacters {
		middle := (rowMax + rowMin) / 2
		if string(character) == "B" {
			rowMin = middle + 1
		} else if string(character) == "F" {
			rowMax = middle
		}
	}
	return rowMax
}

func getColumn(line string) int {
	columnCharacters := line[7:len(line)]
	columnMax := 7
	columnMin := 0
	for _, character := range columnCharacters {
		middle := (columnMax + columnMin) / 2
		if string(character) == "R" {
			columnMin = middle + 1
		} else if string(character) == "L" {
			columnMax = middle
		}
	}
	return columnMax
}

type Position struct {
	row, column int
}

func mapLinesToPlaneRowColumn(lines []string) []Position {
	result := []Position{}
	for _, line := range lines {
		row := getRow(line)
		column := getColumn(line)
		position := Position{row, column}
		result = append(result, position)
	}
	return result
}

func getRowId(position Position) int {
	return position.row * 8 + position.column
}

func mapPositionToRowId(positions []Position) []int {
	rowIds := []int{}
	for _, position := range positions {
		rowId := getRowId(position)
		rowIds = append(rowIds, rowId)
	}
	return rowIds
}

func getMax(rowIds []int) int {
	max := 0
	for _, rowId := range rowIds {
		if rowId > max {
			max = rowId
		}
	}
	return max
}

func getMaximumRowId(filePath string) int {
	text := readFile(filePath)
	lines := strings.Split(text, "\n")
	removeLastLine := lines[:len(lines)-1]
	positions := mapLinesToPlaneRowColumn(removeLastLine)
	rowIds := mapPositionToRowId(positions)
	maxRowId := getMax(rowIds)
	return maxRowId
}

// lazy not returning, wanted to get a good look and the answer was obvious
func printOpenRows(filePath string) {
	text := readFile(filePath)
	lines := strings.Split(text, "\n")
	removeLastLine := lines[:len(lines)-1]
	positions := mapLinesToPlaneRowColumn(removeLastLine)
	rowIds := mapPositionToRowId(positions)
	maxRowId := getMax(rowIds)
	findMissingRows(rowIds, maxRowId)
}

func findMissingRows(rowIds []int, maxRowId int) {
	for i:= 0; i < maxRowId; i++ {
		found := false
		for _, rowId := range rowIds {
			if rowId == i {
				found = true
			}
		}
		if found == false {
			fmt.Println(i)
		}
	}
}

func main() {
	filePath := "./input.txt"
	maxRowId := getMaximumRowId(filePath)
	printOpenRows(filePath)
	fmt.Println(maxRowId)
}
