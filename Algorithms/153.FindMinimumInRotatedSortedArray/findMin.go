package findminimuminrotatedsortedarrayii

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		middle := (start + end) / 2
		if nums[middle] < nums[end] {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return nums[start]
}
