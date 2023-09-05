package findminimuminrotatedsortedarrayii

import "testing"

func Test_search(t *testing.T) {
	testCases := []struct {
		nums         []int
		exceptResult int
	}{
		{
			nums:         []int{1, 3, 5},
			exceptResult: 1,
		},
		{
			nums:         []int{2, 2, 2, 0, 1},
			exceptResult: 0,
		},
		{
			nums:         []int{1, 0, 1, 1, 1},
			exceptResult: 0,
		},
	}

	for _, testCase := range testCases {
		actual := findMin(testCase.nums)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %v, got %v, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
