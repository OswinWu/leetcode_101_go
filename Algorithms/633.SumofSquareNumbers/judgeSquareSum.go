package sumofsquarenumbers

import "math"

func judgeSquareSum(c int) bool {
	right := int(math.Sqrt(float64(c)))
	left := 0
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}
