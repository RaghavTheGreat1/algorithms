package main

import (
	"fmt"
	"math"
)

func main() {
	primeNumbers := getPrimeNumbers(100)
	fmt.Println("Prime Numbers are: ")
	fmt.Println(primeNumbers)
	fmt.Println("Count: ", len(primeNumbers))
}

func getPrimeNumbers(n int) []int {

	numbersList := []int{}

	for i := 0; i <= n; i++ {
		numbersList = append(numbersList, i)
	}
	root := int(math.Sqrt(float64(n)))
	for i := 2; i <= root; i++ {
		for j := i + 1; j <= n; j++ {
			currentNumber := numbersList[j]
			if currentNumber == -1 {
				continue
			}
			if currentNumber%i == 0 {
				numbersList[j] = -1
			}
		}
	}

	primeNumbersList := []int{}

	for i := 0; i <= n; i++ {
		currentNumber := numbersList[i]

		if currentNumber != -1 {
			primeNumbersList =
				append(primeNumbersList, currentNumber)
		}
	}

	// Since 0 & 1 are not prime numbers
	primeNumbersList = primeNumbersList[2:]

	return primeNumbersList
}
