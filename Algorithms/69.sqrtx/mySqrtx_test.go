package mySqrt

import "testing"

func Test_maxProfit(t *testing.T) {
	testCases := []struct {
		input        int
		exceptResult int
	}{
		{
			input:        8,
			exceptResult: 2,
		},
		{
			input:        4,
			exceptResult: 2,
		},
		{
			input:        11,
			exceptResult: 3,
		},
	}

	for _, testCase := range testCases {
		actual := mySqrt(testCase.input)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
