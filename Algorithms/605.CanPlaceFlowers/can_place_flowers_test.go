package canPlaceFlowers

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	testCases := []struct {
		inputArr     []int
		inputNum     int
		exceptResult bool
	}{
		{
			inputArr:     []int{1, 0, 0, 0, 1},
			inputNum:     1,
			exceptResult: true,
		},
		{
			inputArr:     []int{0, 0, 0, 0},
			inputNum:     3,
			exceptResult: false,
		},
		{
			inputArr:     []int{0},
			inputNum:     1,
			exceptResult: true,
		},
		{
			inputArr:     []int{1, 0, 0, 0, 1},
			inputNum:     2,
			exceptResult: false,
		},
		{
			inputArr:     []int{0, 0, 1, 0, 1},
			inputNum:     1,
			exceptResult: true,
		},
	}

	for _, testCase := range testCases {
		actual := canPlaceFlowers(testCase.inputArr, testCase.inputNum)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %v, got %v, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
