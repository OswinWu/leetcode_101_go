package sumofsquarenumbers

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	testCases := []struct {
		c    int
		want bool
	}{
		{5, true},
		{3, false},
	}
	for _, tc := range testCases {
		if got := judgeSquareSum(tc.c); got != tc.want {
			t.Errorf("judgeSquareSum(%v) = %v, want %v", tc.c, got, tc.want)
		}
	}
}
