package main

import (
	"fmt"
	"strings"
)

func main() {
	romantext := "MCMXCIV"

	fmt.Println("romanInt : ", NewTesr(romantext))

}

func NewTesr(romantext string) int {
	var romanInt int
	entered_here := 0
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	romantextArray := strings.Split(romantext, "")

	for i := 0; i < len(romantextArray); i++ {
		if i < len(romantextArray)-1 {
			if romantextArray[i] == "I" && romantextArray[i+1] == "V" || romantextArray[i] == "I" && romantextArray[i+1] == "X" {
				romanInt = romanInt + (roman[romantextArray[i+1]] - roman[romantextArray[i]])
				entered_here++
			} else if romantextArray[i] == "X" && romantextArray[i+1] == "L" || romantextArray[i] == "X" && romantextArray[i+1] == "C" {
				romanInt = romanInt + roman[romantextArray[i+1]] - roman[romantextArray[i]]
				entered_here++
			} else if romantextArray[i] == "C" && romantextArray[i+1] == "D" || romantextArray[i] == "C" && romantextArray[i+1] == "M" {
				romanInt = romanInt + roman[romantextArray[i+1]] - roman[romantextArray[i]]
				entered_here++
			}
		}
		if entered_here == 0 {
			romanInt = romanInt + roman[romantextArray[i]]
		}

	}

	return romanInt
}
