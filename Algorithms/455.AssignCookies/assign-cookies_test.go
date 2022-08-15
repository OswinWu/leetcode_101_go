package assigncookies

import (
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	testCases := []struct {
		inputChildren []int
		inputCookies  []int
		excepetResult int
	}{
		{
			inputChildren: []int{1, 2},
			inputCookies:  []int{1, 2, 3},
			excepetResult: 2,
		},
		{
			inputChildren: []int{3, 2, 1},
			inputCookies:  []int{1, 1, 1},
			excepetResult: 1,
		},
		{
			inputChildren: []int{10, 9, 8, 7},
			inputCookies:  []int{5, 6, 7, 8},
			excepetResult: 2,
		},
	}

	for _, testCase := range testCases {
		actual := findContentChildren(testCase.inputChildren, testCase.inputCookies)
		if actual != testCase.excepetResult {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.excepetResult, actual, testCase)
		}
	}
}
