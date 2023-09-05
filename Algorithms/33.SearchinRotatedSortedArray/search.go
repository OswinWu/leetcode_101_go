package searchinrotatedsortedarray

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[start] <= nums[middle] {
			if target >= nums[start] && target < nums[middle] {
				end = middle - 1
			} else {
				start = middle + 1
			}
		} else {
			if target > nums[middle] && target <= nums[end] {
				start = middle + 1
			} else {
				end = middle - 1
			}
		}
	}
	return -1
}
