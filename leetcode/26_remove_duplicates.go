package main

func RemoveDuplicates(nums []int) int {
	// Get the length of the input slice
	n := len(nums)

	// Create a map to store non-duplicate numbers
	nonDuplicates := make(map[int]int)

	// Initialize variables
	var currentNumber int
	newIndex := 0

	// Iterate over the input slice
	for i := 0; i < n; i++ {
		currentNumber = nums[i]

		// Check if the current number is already in the nonDuplicates map
		if nonDuplicates[currentNumber] > 0 {
			// If it is, continue to the next iteration without updating the slice or map
			continue
		} else {
			// If it's not, update the slice and map with the current number
			nums[newIndex] = currentNumber
			nonDuplicates[currentNumber] = 1
			newIndex++
		}
	}

	// Return the new length of the slice after removing duplicates
	return newIndex
}
