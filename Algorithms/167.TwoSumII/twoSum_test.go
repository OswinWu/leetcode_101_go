package twosumii

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_towSum(t *testing.T) {
	testCases := []struct {
		numbers []int
		target  int
		result  []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, %v", tc.numbers, tc.target), func(t *testing.T) {
			if result := twoSum(tc.numbers, tc.target); !reflect.DeepEqual(result, tc.result) {
				t.Errorf("expected %v, got %v", tc.result, result)
			}
		})
	}

}
