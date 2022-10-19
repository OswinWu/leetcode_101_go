package countStudents

import "testing"

func Test_countStudents(t *testing.T) {
	testCases := []struct {
		students     []int
		sandwiches   []int
		exceptResult int
	}{
		{
			students:     []int{1, 1, 0, 0},
			sandwiches:   []int{0, 1, 0, 1},
			exceptResult: 0,
		},
		{
			students:     []int{1, 1, 1, 0, 0, 1},
			sandwiches:   []int{1, 0, 0, 0, 1, 1},
			exceptResult: 3,
		},
		{
			students:     []int{1, 1},
			sandwiches:   []int{0, 1},
			exceptResult: 2,
		},
	}

	for _, testCase := range testCases {
		actual := countStudents(testCase.students, testCase.sandwiches)
		if actual != testCase.exceptResult {
			t.Errorf("error result, except %d, got %d, testCase %v", testCase.exceptResult, actual, testCase)
		}
	}
}
