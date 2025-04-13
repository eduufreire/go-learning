package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	numbersRoman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if len(s) < 1 || len(s) > 15 {
		return 0
	}

	splitText := strings.Split(s, "")
	sum := 0
	jumpNext := false

	for key, value := range splitText {

		if jumpNext {
			jumpNext = false
			continue
		}

		letterToInt := numbersRoman[value]

		if key < len(splitText) - 1 {
			if value == "I" && (splitText[key+1] == "V" || splitText[key+1] == "X") {
				letterToInt = numbersRoman[splitText[key+1]]
				sum = sum + (letterToInt - 1)
				jumpNext = true
				continue
			}
	
			if value == "X" && ( splitText[key+1] == "L" || splitText[key+1] == "C") {
				letterToInt = numbersRoman[splitText[key+1]]
				sum = sum + (letterToInt - 10)
				jumpNext = true
				continue
			}
	
			if value == "C" &&(splitText[key+1] == "D" || splitText[key+1] == "M") {
				letterToInt = numbersRoman[splitText[key+1]]
				sum = sum + (letterToInt - 100)
				jumpNext = true
				continue
			}
		}
		
		sum += letterToInt
	}
	return sum
}
