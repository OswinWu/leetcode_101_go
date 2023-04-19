package nondecreasingarray

func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	flag := false

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if i == 0 {
				nums[i] = nums[i+1]
			} else if i == len(nums)-2 {
				nums[i+1] = nums[i]
			} else {
				if nums[i-1] <= nums[i+1] {
					nums[i] = nums[i-1]
				} else {
					nums[i+1] = nums[i]
				}
			}
			flag = true
			break
		}
	}

	if flag {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				return false
			}
		}
	}

	return true
}
