package nonOverLappingIntervals

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	pre := intervals[0][1]
	removedNum := 0
	size := len(intervals)
	index := 1
	for {
		if index == size {
			break
		}
		if pre > intervals[index][0] {
			removedNum++
		} else {
			pre = intervals[index][1]
		}
		index++
	}
	return removedNum
}
