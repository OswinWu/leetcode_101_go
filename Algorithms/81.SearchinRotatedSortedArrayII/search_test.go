package searchinrotatedsortedarrayii

import "testing"

func Test_search(t *testing.T) {
	testCases := []struct {
		nums         []int
		target       int
		exceptResult bool
	}{
		{
			nums:         []int{1, 0, 1, 1, 1},
			target:       0,
			exceptResult: true,
		},
		{
			nums:         []int{1, 3, 1, 1},
			target:       0,
			exceptResult: false,
		},
	}

	for _, testCase := range testCases {
		actual := search(testCase.nums, testCase.target)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %v, got %v, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
