package main

func ContainsDuplicate(nums []int) bool {

	visitedNumbers := make(map[int]int)

	for _, v := range nums {
		visitedNumbers[v] += 1
		if visitedNumbers[v] > 1 {
			return true
		}
	}

	return false
}
