package main

import (
	"fmt"
	"math"
)

func main() {
	// Calling the getPrimeNumbers function with input 23
	primeNumbers := getPrimeNumbers(100)

	// Printing the list of prime numbers
	fmt.Println("Prime Numbers are: ")
	fmt.Println(primeNumbers)

	// Printing the count of prime numbers
	fmt.Println("Count: ", len(primeNumbers))
}

// getPrimeNumbers generates a list of prime numbers from range [0,n] using Sieve of Eratosthenes algorithm.
// It takes in an integer n as input and returns a slice of prime numbers.
func getPrimeNumbers(n int) []int {
	// Creating a list of numbers from 0 to n
	numbersList := []int{}
	for i := 0; i <= n; i++ {
		numbersList = append(numbersList, i)
	}

	// Finding the square root of n for optimization
	root := int(math.Sqrt(float64(n)))

	// Sieve of Eratosthenes algorithm
	// Marking multiples of numbers as -1
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

	// Extracting prime numbers from the list
	primeNumbersList := []int{}
	for i := 0; i <= n; i++ {
		currentNumber := numbersList[i]
		if currentNumber != -1 {
			primeNumbersList = append(primeNumbersList, currentNumber)
		}
	}

	// Since 0 and 1 are not prime numbers, removing them from the list
	primeNumbersList = primeNumbersList[2:]

	return primeNumbersList
}
