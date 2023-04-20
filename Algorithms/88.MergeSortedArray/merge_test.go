package mergesortedarray

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	testCases := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		target []int
	}{
		{
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			target: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			target: []int{1},
		},
		{
			nums1:  []int{0},
			m:      0,
			nums2:  []int{1},
			n:      1,
			target: []int{1},
		},
		{
			nums1:  []int{2, 0},
			m:      1,
			nums2:  []int{1},
			n:      1,
			target: []int{1, 2},
		},
		{
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{4, 5, 6},
			n:      3,
			target: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, testCase := range testCases {
		merge(testCase.nums1, testCase.m, testCase.nums2, testCase.n)
		if !reflect.DeepEqual(testCase.nums1, testCase.target) {
			t.Errorf("error result, except %v, got %v, testCase %v", testCase.target, testCase.nums1, testCase)
		}
	}
}
