package main

func SearchInsert(nums []int, target int) int {
	n := len(nums)
	low, high := 0, n-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
