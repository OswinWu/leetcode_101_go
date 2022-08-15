package candy

import "testing"

func Test_candy(t *testing.T) {
	testCases := []struct {
		input         []int
		excepetResult int
	}{
		{
			input:         []int{1, 0, 2},
			excepetResult: 5,
		},
		{
			input:         []int{1, 2, 2},
			excepetResult: 4,
		},
		{
			input:         []int{1, 3, 2, 2, 1},
			excepetResult: 7,
		},
		{
			input:         []int{1, 2, 87, 87, 87, 2, 1},
			excepetResult: 13,
		},
	}

	for _, testCase := range testCases {
		actual := candy(testCase.input)
		if actual != testCase.excepetResult {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.excepetResult, actual, testCase)
		}
	}
}
