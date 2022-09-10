package findMinArrowShots

import "sort"

func findMinArrowShots(points [][]int) int {
	pointLen := len(points)
	groupNum := pointLen
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	right := points[0][1]
	for i := 1; i < pointLen; i++ {
		if right >= points[i][0] {
			groupNum = groupNum - 1
		} else {
			right = points[i][1]
		}
	}
	return groupNum
}
