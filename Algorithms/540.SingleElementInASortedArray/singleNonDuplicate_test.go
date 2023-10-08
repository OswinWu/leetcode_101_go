package singleelementinasortedarray

import "testing"

func Test_search(t *testing.T) {
	testCases := []struct {
		nums         []int
		exceptResult int
	}{
		{
			nums:         []int{1, 1, 2},
			exceptResult: 2,
		},
		{
			nums:         []int{1, 1, 2, 2, 3},
			exceptResult: 3,
		},
		{
			nums:         []int{7, 7, 10, 11, 11, 12, 12},
			exceptResult: 10,
		},
		{
			nums:         []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			exceptResult: 2,
		},
		{
			nums:         []int{1, 1, 2, 2, 3, 3, 4},
			exceptResult: 4,
		},
		{
			nums:         []int{3, 3, 7, 7, 10, 11, 11},
			exceptResult: 10,
		},
		{
			nums:         []int{1, 1, 2, 2, 3, 3, 4, 4, 8, 8, 0},
			exceptResult: 0,
		},
	}

	for _, testCase := range testCases {
		actual := singleNonDuplicate(testCase.nums)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %v, got %v, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
