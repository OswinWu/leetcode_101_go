package longestwordindictionarythroughdeletin

import (
	"testing"
)

func Test_findLongestWord(t *testing.T) {
	testCases := []struct {
		s    string
		dict []string
		want string
	}{
		{
			"abpcplea",
			[]string{"ale", "apple", "monkey", "plea"},
			"apple",
		},
		{
			"abpcplea",
			[]string{"a", "b", "c"},
			"a",
		},
		{
			"bab",
			[]string{"ba", "ab", "a", "b"},
			"ab",
		},
	}

	for _, tc := range testCases {
		got := findLongestWord(tc.s, tc.dict)
		if got != tc.want {
			t.Errorf("Input: %v, %v, want: %v, got: %v", tc.s, tc.dict, tc.want, got)
		}
	}
}
