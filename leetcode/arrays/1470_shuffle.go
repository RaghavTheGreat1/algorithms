package arrays

func Shuffle(nums []int, n int) []int {
	newArray := []int{}

	for i := 0; i < n; i++ {
		newArray = append(newArray, nums[i], nums[i+n])
	}

	return newArray
}
