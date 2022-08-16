package assignCookies

import (
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	testCases := []struct {
		inputChildren []int
		inputCookies  []int
		exceptResult  int
	}{
		{
			inputChildren: []int{1, 2},
			inputCookies:  []int{1, 2, 3},
			exceptResult:  2,
		},
		{
			inputChildren: []int{3, 2, 1},
			inputCookies:  []int{1, 1, 1},
			exceptResult:  1,
		},
		{
			inputChildren: []int{10, 9, 8, 7},
			inputCookies:  []int{5, 6, 7, 8},
			exceptResult:  2,
		},
	}

	for _, testCase := range testCases {
		actual := findContentChildren(testCase.inputChildren, testCase.inputCookies)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
