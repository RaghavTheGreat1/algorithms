package main

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	mergedSortedArray := make([]int, m+n)

	i, j, k := 0, 0, 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			mergedSortedArray[k] = nums1[i]
			i++
		} else {
			mergedSortedArray[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		mergedSortedArray[k] = nums1[i]
		i++
		k++
	}

	for j < n {
		mergedSortedArray[k] = nums2[j]
		j++
		k++
	}

	return mergedSortedArray
}
