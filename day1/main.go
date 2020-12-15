package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
	"strconv"
)

func readFile(filePath string) string {
	contents, _ := ioutil.ReadFile(filePath)
	var s string = string(contents)
	return s
}


func convertToNumbers(toConvert []string) []int {
    var toPlace = []int{}

    for _, stringNumber := range toConvert {
        number, _ := strconv.Atoi(stringNumber)
        toPlace = append(toPlace, number)
    }
	return toPlace
}

type Pair struct {
    a, b, multiplied int
}

func getPair(numbers []int) Pair {
	// note that numbers are sorted
	for i, number1 := range numbers {
		for j := i+1; j < len(numbers); j++ {
			number2 := numbers[j]
			if(number1 + number2 == 2020){
				return Pair{number1, number2, number1*number2}
			}
		}
	}
	panic("There is no solution")
}

func solveDay1(filePath string) Pair {
	text := readFile(filePath)
	lines := strings.Split(text,"\n")
	numbers := convertToNumbers(lines)
	sort.Ints(numbers) // in place
	pair := getPair(numbers)
	return pair
}

func main(){
	filePath := "./input.txt"
	pair := solveDay1(filePath)
	fmt.Println(pair)
}
