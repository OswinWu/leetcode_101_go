package nonOverLappingIntervals

import (
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	testCases := []struct {
		input        [][]int
		exceptResult int
	}{
		{
			input: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{1, 3},
			},
			exceptResult: 1,
		},
		{
			input: [][]int{
				{1, 2},
				{1, 2},
				{1, 2},
			},
			exceptResult: 2,
		},
		{
			input: [][]int{
				{1, 2},
				{2, 3},
			},
			exceptResult: 0,
		},
		{
			input: [][]int{
				{1, 100},
				{11, 22},
				{1, 11},
				{2, 12},
			},
			exceptResult: 2,
		},
	}

	for _, testCase := range testCases {
		actual := eraseOverlapIntervals(testCase.input)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
