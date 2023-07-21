package main

import (
	"fmt"
)

func main() {
	items := []int{9, 5, 7, 3, 1, 7, 2, 10, 4, 8}
	fmt.Println("before:", items)
	result := mergeSort(items)
	fmt.Println("after : ", result)
}

func mergeSort(items []int) []int {
	// Base case: if the input slice has less than 2 elements, return it as is
	if len(items) < 2 {
		return items
	}

	// Divide the input slice into two halves
	middle := len(items) / 2
	left := mergeSort(items[:middle])
	right := mergeSort(items[middle:])

	// Merge the sorted left and right halves
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	i := 0
	j := 0

	// Compare the elements from both halves and append them to the result slice in sorted order
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements from the left and right halves
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
