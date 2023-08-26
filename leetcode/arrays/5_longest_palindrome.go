package arrays

func LongestPalindromeBruteForce(s string) string {
	n := len(s)

	largestLength := 0
	largestPalindrome := ""
	for i := 0; i < n; i++{
		for j:= i; j<n; j++ {
			slice := s[i:j+1]
			isPalindrome := checkPalindrome(slice)
			if isPalindrome{
				currentLength := j-i + 1
				if currentLength > largestLength{
					largestPalindrome = slice
					largestLength = currentLength
				}
			}
		}
	}
	return largestPalindrome
}

func checkPalindrome(s string) bool{
	n := len(s)
	for i := 0; i < n; i++{
		if s[i] != s[n-i-1]{
			return false
		} 
	}
	return true
}