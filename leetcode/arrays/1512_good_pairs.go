package arrays

func NumIdenticalPairs(nums []int) int {
	count := 0

	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
	}

	return count
}
