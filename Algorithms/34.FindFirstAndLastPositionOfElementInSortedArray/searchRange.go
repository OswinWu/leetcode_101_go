package searchRange

func searchRange(nums []int, target int) []int {
	right := len(nums)
	if right == 0 {
		return []int{-1, -1}
	}

	if right == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	if right == 2 {
		if nums[0] == target && nums[1] == target {
			return []int{0, 1}
		}
		if nums[0] == target {
			return []int{0, 0}
		}
		if nums[1] == target {
			return []int{1, 1}
		}
		return []int{-1, -1}
	}

	left := 0
	middle := (left + right) / 2
	leftResult := searchRange(nums[left:middle], target)
	rightResult := searchRange(nums[middle:right], target)
	if leftResult[0] != -1 && rightResult[0] != -1 {
		return []int{leftResult[0], middle + rightResult[1]}
	}
	if leftResult[0] != -1 {
		return leftResult
	}
	if rightResult[0] != -1 {
		return []int{rightResult[0] + middle, rightResult[1] + middle}
	}
	return []int{-1, -1}
}
