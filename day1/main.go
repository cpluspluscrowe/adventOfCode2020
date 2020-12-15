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
    a, b int
}

func getMagicNumber(numbers []int) Pair {
	// note that numbers are sorted
	for i, number1 := range numbers {
		for j := i+1; j < len(numbers); j++ {
			number2 := numbers[j]
			if(number1 + number2 == 2020){
				if(number1 * number2 == 514579){
					return Pair{number1, number2}
				}
			}
		}
	}
	panic("There is no solution")
}

func main(){
	filePath := "./input.txt"
	text := readFile(filePath)
	lines := strings.Split(text,"\n")
	numbers := convertToNumbers(lines)
	sort.Ints(numbers) // in place
	magicNumber := getMagicNumber(numbers)
	fmt.Println(magicNumber)
}
