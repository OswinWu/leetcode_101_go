package findMinArrowShots

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	testCases := []struct {
		inputArr     [][]int
		exceptResult int
	}{
		{
			inputArr: [][]int{
				{10, 16},
				{2, 8},
				{1, 6},
				{7, 12},
			},
			exceptResult: 2,
		},
		{
			inputArr: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
			},
			exceptResult: 4,
		},
		{
			inputArr: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
			},
			exceptResult: 2,
		},
		// {
		// 	inputArr: [][]int{
		// 		{1, 2},
		// 		{2, 3},
		// 		{3, 4},
		// 		{1, 3},
		// 	},
		// 	exceptResult: 1,
		// },
		// {
		// 	inputArr: [][]int{
		// 		{1, 2},
		// 		{2, 3},
		// 		{3, 4},
		// 		{1, 3},
		// 	},
		// 	exceptResult: 1,
		// },
	}

	for _, testCase := range testCases {
		actual := findMinArrowShots(testCase.inputArr)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %v, got %v, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
