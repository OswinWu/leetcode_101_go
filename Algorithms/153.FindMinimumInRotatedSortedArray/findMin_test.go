package findminimuminrotatedsortedarrayii

import "testing"

func Test_search(t *testing.T) {
	testCases := []struct {
		nums         []int
		exceptResult int
	}{
		{
			nums:         []int{3, 4, 5, 1, 2},
			exceptResult: 1,
		},
		{
			nums:         []int{4, 5, 6, 7, 0, 1, 2},
			exceptResult: 0,
		},
		{
			nums:         []int{11, 13, 15, 17},
			exceptResult: 11,
		},
	}

	for _, testCase := range testCases {
		actual := findMin(testCase.nums)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %v, got %v, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
