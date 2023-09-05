package searchinrotatedsortedarray

import "testing"

func Test_search(t *testing.T) {
	testCases := []struct {
		nums         []int
		target       int
		exceptResult int
	}{
		{
			nums:         []int{4, 5, 6, 7, 0, 1, 2},
			target:       0,
			exceptResult: 4,
		},
		{
			nums:         []int{4, 5, 6, 7, 0, 1, 2},
			target:       3,
			exceptResult: -1,
		},
	}

	for _, testCase := range testCases {
		actual := search(testCase.nums, testCase.target)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %v, got %v, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
