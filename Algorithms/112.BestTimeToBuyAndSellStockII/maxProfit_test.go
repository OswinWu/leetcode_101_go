package maxProfit

import "testing"

func Test_maxProfit(t *testing.T) {
	testCases := []struct {
		input        []int
		exceptResult int
	}{
		{
			input:        []int{7, 1, 5, 3, 6, 4},
			exceptResult: 7,
		},
		{
			input:        []int{1, 2, 3, 4, 5},
			exceptResult: 4,
		},
		{
			input:        []int{7, 6, 4, 1},
			exceptResult: 0,
		},
	}

	for _, testCase := range testCases {
		actual := maxProfit(testCase.input)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
