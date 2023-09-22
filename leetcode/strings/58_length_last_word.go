package strings

func LengthOfLastWord(s string) int {

    count := 0
    for i := len(s)-1; i >= 0; i--{
        currentChar := string(s[i])
        if currentChar != " "{
            count++
        }

        if count > 0 && currentChar == " "{
            return count
        }
    }
    return count
}
