package arrays

func BuildArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i, v := range nums {
		ans[i] = nums[v]
	}

	return ans
}
