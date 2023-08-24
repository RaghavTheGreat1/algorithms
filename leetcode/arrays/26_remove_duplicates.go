package arrays

func RemoveDuplicates(nums []int) int {

	n := len(nums)

	visitedNumbers := make(map[int]int)

	var currentNumber int
	newIndex := 0

	for i := 0; i < n; i++ {
		currentNumber = nums[i]

		if visitedNumbers[currentNumber] > 0 {

			continue
		} else {

			nums[newIndex] = currentNumber
			visitedNumbers[currentNumber] = 1
			newIndex++
		}
	}

	return newIndex
}
