package validpalindromeii

func validPalindrome(s string) bool {
	left := 0
	right := len(s)
	if len(s) < 2 {
		return true
	}
	for left < right {
		if s[left] != s[right-1] {
			return isPalindrome(s[left:right-1]) || isPalindrome(s[left+1:right])
		}
		left++
		right--
	}

	return true
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s)
	if len(s) < 2 {
		return true
	}
	for left < right {
		if s[left] != s[right-1] {
			return false
		}
		left++
		right--
	}

	return true
}
