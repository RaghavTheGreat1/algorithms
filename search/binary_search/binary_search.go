package main

import "fmt"

func main() {
	// Numbers should be sorted in ascending order
	numbers := []int{1, 2, 5, 8, 10, 12, 90}

	index := binarySearch(numbers, 1)

	if index == -1 {
		fmt.Println("No element found")
	} else {
		fmt.Println("Element found at index: ", index)
	}
}

// binarySearch performs a binary search algorithm to find the index of the key element in the given slice of numbers.
func binarySearch(sortedNumbers []int, key int) int {
	n := len(sortedNumbers)

	low := 0
	high := n

	// Call binarySearchNonRecursive to perform a non-recursive binary search
	index := binarySearchNonRecursive(sortedNumbers, low, high, key)

	// Call binarySearchRecursive to perform a recursive binary search
	index = binarySearchRecursive(sortedNumbers, low, high, key)

	return index
}

// binarySearchNonRecursive performs a non-recursive binary search on the given slice of numbers.
func binarySearchNonRecursive(sortedNumbers []int, low int, high int, key int) int {
	lowIndex := low
	highIndex := high

	for lowIndex < highIndex {
		mid := (lowIndex + highIndex) / 2

		// Check if the mid element is equal to the key
		if key == sortedNumbers[mid] {
			return mid
			// If key is greater than mid element, then key is to the right half of the array
			// hence divide the array by altering the lower index
		} else if key > sortedNumbers[mid] {
			lowIndex = mid + 1
			// Here, key is less than the mid element, hence the key is to the left half of the array
		} else {
			highIndex = mid - 1
		}
	}

	// If the key element is not found, return -1
	return -1
}

// binarySearchRecursive performs a recursive binary search on the given slice of numbers.
func binarySearchRecursive(sortedNumbers []int, low int, high int, key int) int {
	mid := (low + high) / 2

	// Check if the lower index is less than the higher index
	if low < high {
		// Check if the mid element is equal to the key
		if sortedNumbers[mid] == key {
			return mid
			// If key is greater than mid element, then key is to the right half of the array
			// hence divide the array by altering the lower index
		} else if key > sortedNumbers[mid] {
			return binarySearchRecursive(sortedNumbers, mid+1, high, key)
			// Here, key is less than the mid element, hence the key is to the left half of the array
		} else {
			return binarySearchRecursive(sortedNumbers, low, mid-1, key)
		}
	} else {
		// This case reaches if the key is at the ends of the array & when program reaches to the last recursion or smallest array of size 1
		if sortedNumbers[mid] == key {
			return mid
		} else {
			return -1
		}
	}
}
