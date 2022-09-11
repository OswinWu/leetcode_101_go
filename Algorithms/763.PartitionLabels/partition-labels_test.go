package partitionLabels

import "testing"

func Test_partitionLabels(t *testing.T) {
	testCases := []struct {
		input        string
		exceptResult []int
	}{
		{
			input:        "ababcbacadefegdehijhklij",
			exceptResult: []int{9, 7, 8},
		},
		{
			input:        "eccbbbbdec",
			exceptResult: []int{10},
		},
		// {
		// 	input:        ,
		// 	exceptResult: []int{},
		// },
		// {
		// 	input:        ,
		// 	exceptResult: []int{},
		// },
	}

	for _, testCase := range testCases {
		actual := partitionLabels(testCase.input)
		if len(actual) != len(testCase.exceptResult) {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.exceptResult, actual, testCase)
		} else {
			for i := 0; i < len(actual); i++ {
				if actual[i] != testCase.exceptResult[i] {
					t.Errorf("error result, except %d, got %d, testCase %v", testCase.exceptResult, actual, testCase)
					break
				}
			}
		}
	}
}
