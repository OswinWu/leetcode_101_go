package validpalindromeii

import "testing"

func Test_validPalindrome(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
	}{
		{"aba", true},
		{"abca", true},
		{"abc", false},
		{"ab", true},
		{"a", true},
		{"", true},
		{"abccba", true},
	}
	for _, tc := range testCases {
		if got := validPalindrome(tc.s); got != tc.want {
			t.Errorf("validPalindrome(%v) = %v, want %v", tc.s, got, tc.want)
		}
	}
}
