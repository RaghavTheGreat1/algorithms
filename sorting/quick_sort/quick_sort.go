package main

import "fmt"

func main() {
	// Input array
	arr := []int{7, 2, 1, 6, 8, 5, 3, 4}

	// Sort the array
	sortedArr := quickSort(arr)

	// Print the sorted array
	fmt.Println(sortedArr)
}

// Quicksort function
func quickSort(arr []int) []int {
	// Base case: an empty or single-element array is already sorted
	if len(arr) < 2 {
		return arr
	}

	// Choose a pivot element (for simplicity, we'll use the first element)
	pivot := arr[0]

	// Initialize the left and right partitions
	left := []int{}
	right := []int{}

	// Partition the array
	for _, num := range arr[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	// Recursively sort the left and right partitions
	left = quickSort(left)
	right = quickSort(right)

	// Concatenate the sorted partitions and the pivot element
	return append(append(left, pivot), right...)
}
