package mySqrt

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left := 1
	right := x
	for {
		middle := (left + right) / 2
		result := middle * middle
		if result == x {
			return middle
		} else if result > x {
			right = middle
		} else {
			left = middle
		}
		if (right - left) == 1 {
			return left
		}
	}
}
