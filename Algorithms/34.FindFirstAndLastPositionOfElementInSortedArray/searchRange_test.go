package searchRange

import "testing"

func Test_searchRange(t *testing.T) {
	testCases := []struct {
		nums         []int
		target       int
		exceptResult []int
	}{
		{
			nums:         []int{5, 7, 7, 8, 8, 10},
			target:       8,
			exceptResult: []int{3, 4},
		},
		{
			nums:         []int{1},
			target:       0,
			exceptResult: []int{-1, -1},
		},
		{
			nums:         []int{5, 7, 7, 8, 8, 10},
			target:       6,
			exceptResult: []int{-1, -1},
		},
	}

	for _, testCase := range testCases {
		actual := searchRange(testCase.nums, testCase.target)
		if actual[0] != testCase.exceptResult[0] || actual[1] != testCase.exceptResult[1] {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
