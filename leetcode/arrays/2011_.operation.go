package arrays

func FinalValueAfterOperations(operations []string) int {
	x := 0
	for _, v := range operations {
		if v[0:2] == "++" || v[1:] == "++" {
			x++
		} else {
			x--
		}
	}
	return x
}
