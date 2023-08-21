package main

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	nonDups := make(map[int]int)

    var number int
    j := 0
	for i:= 0; i<n; i++{
        number = nums[i]
		if nonDups[number] > 0{
            continue
        }else{
            nums[j]=number
            nonDups[number] = 1
            j++
        }
	}	

	return j
}