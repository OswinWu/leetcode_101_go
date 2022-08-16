package candy

import "testing"

func Test_candy(t *testing.T) {
	testCases := []struct {
		input        []int
		exceptResult int
	}{
		{
			input:        []int{1, 0, 2},
			exceptResult: 5,
		},
		{
			input:        []int{1, 2, 2},
			exceptResult: 4,
		},
		{
			input:        []int{1, 3, 2, 2, 1},
			exceptResult: 7,
		},
		{
			input:        []int{1, 2, 87, 87, 87, 2, 1},
			exceptResult: 13,
		},
	}

	for _, testCase := range testCases {
		actual := candy(testCase.input)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
