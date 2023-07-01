package main

import "fmt"

func main() {
	// Call the matchString function to check if the text contains the subString
	contains := matchString("Hello World! This is Go", "Go")

	// Print the result
	fmt.Println(contains)
}

// matchString checks if the given text contains the subString
func matchString(text string, subString string) bool {
	// Get the lengths of the text and subString
	n := len(text)
	m := len(subString)

	// Iterate through the text to check for the subString
	for i := 0; i < n; i++ {
		// Initialize a variable to count the matching characters
		matchCount := 0

		// Iterate through the subString and compare it to the text
		for j := 0; j < m; j++ {
			// If the characters don't match, break out of the loop
			if subString[j] != text[i+j] {
				break
			}
			// If the characters match, increment the matchCount
			matchCount++
		}

		// If the matchCount is equal to the length of the subString,
		// it means the subString has been found in the text
		if matchCount == m {
			return true
		}
	}

	// If the subString is not found in the text, return false
	return false
}