package main

import "fmt"

func main() {

	number := romanToInt("I")

	fmt.Println(number)
}

func romanToInt(s string) int {

	romanAsInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sLen := len(s)

	var number int = 0
	i := 0
	for i < sLen {
		currentRomanEq := romanAsInt[s[i:i+1]]

		var nextRomanEq int
		if i == sLen-1 {
			nextRomanEq = 0
		} else {
			nextRomanEq = romanAsInt[s[i+1:i+2]]
		}

		if currentRomanEq >= nextRomanEq {
			number += currentRomanEq
			i++
		} else if currentRomanEq < nextRomanEq {
			number += nextRomanEq - currentRomanEq
			i += 2
		}
	}

	return number
}
