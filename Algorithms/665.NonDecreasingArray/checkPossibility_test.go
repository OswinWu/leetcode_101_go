package nondecreasingarray

import "testing"

func Test_checkPossibility(t *testing.T) {
	testCases := []struct {
		input        []int
		exceptResult bool
	}{
		{
			input:        []int{4, 2, 3},
			exceptResult: true,
		},
		{
			input:        []int{4, 2, 1},
			exceptResult: false,
		},
	}
	for _, testCase := range testCases {
		actual := checkPossibility(testCase.input)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %t, got %t, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
