package main

func PlusOneOptimal(digits []int) []int {
	n := len(digits)
	carry := 1
	for i := n-1; i >= 0; i-- {
		if carry == 0 {
			return digits
		}
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
			carry = 1
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

func PlusOne(digits []int) []int {
	n := len(digits)

	digits[n-1] = digits[n-1]+1
	for i := n-1; i > 0 ; i--{
		if digits[i]!= 10{
			return digits
		}else{
			digits[i] = 0
			digits[i-1] = digits[i-1] + 1
		}
	}

	if digits[0] == 10{
		newDigits := append([]int{1}, digits...)
		newDigits[1] = 0
		return newDigits
	}

	return digits

}