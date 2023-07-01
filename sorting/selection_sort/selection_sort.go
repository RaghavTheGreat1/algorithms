package main

import "fmt"

func main() {
	// Create an unsorted array of integers
	unsortedArray := []int{1, 10, 5, 90, 12, 5, 8, 2}

	// Call the selectionSort function to sort the array
	sortedArray := selectionSort(unsortedArray)

	// Print the sorted array
	fmt.Println(sortedArray)
}

// selectionSort sorts an array of integers using the Selection Sort algorithm
func selectionSort(numbers []int) []int {
	// Create a copy of the input array
	// to avoid modifying the original array
	unsortedNumbers := make([]int, len(numbers))
	copy(unsortedNumbers, numbers)

	// Get the length of the array
	n := len(unsortedNumbers)

	// Iterate over the array
	for i := 0; i < n-1; i++ {
		// Assume the current element is the minimum
		minimumIndex := i

		// Find the minimum element in the unsorted portion of the array
		for j := i + 1; j < n; j++ {
			if unsortedNumbers[j] < unsortedNumbers[minimumIndex] {
				// Found a new minimum element
				minimumIndex = j
			}
		}

		// Swap the current element with the minimum element
		unsortedNumbers[i], unsortedNumbers[minimumIndex] = unsortedNumbers[minimumIndex], unsortedNumbers[i]
	}

	// Return the sorted array
	return unsortedNumbers
}
