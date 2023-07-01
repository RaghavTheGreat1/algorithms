package main

import "fmt"

func main() {
	// Create an array of integers to search from
	numbers := []int{1, 10, 5, 90, 12, 5, 8, 2}

	// Call the linearSearch function to search for the key element 90
	index := linearSearch(numbers, 90)

	// Check if the key element was found or not
	if index == -1 {
		fmt.Println("Element not found")
	} else {
		fmt.Print("Element found at index: ", index)
	}
}

// linearSearch performs a linear search to find the position of a key element in an array
//
// Returns -1 is no element is found
func linearSearch(numbers []int, key int) int {
	// Get the length of the array
	n := len(numbers)

	// Iterate through the array to search for the key element
	for i := 0; i < n; i++ {
		// Check if the current element is equal to the key element
		if key == numbers[i] {
			// Return the index of the key element if found
			return i
		}
	}

	// If the key element is not found, return -1
	return -1
}
