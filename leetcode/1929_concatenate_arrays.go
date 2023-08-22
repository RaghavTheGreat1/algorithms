package main

func GetConcatenationOneLiner(nums []int) []int {
	return append(nums, nums...)
}

func GetConcatenation(nums []int) []int {
	n := len(nums)
	newArray := make([]int, 2*n)

	for i, v := range nums {
		newArray[i], newArray[n+i] = v, v
	}

	return newArray
}
