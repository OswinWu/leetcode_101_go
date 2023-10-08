package singleelementinasortedarray

func singleNonDuplicate(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		if start == end-1 {
			if start == 0 || nums[start] != nums[start-1] {
				return nums[start]
			} else {
				return nums[end]
			}
		}
		mid := (start + end) / 2
		if (nums[mid] != nums[mid-1]) && (nums[mid] != nums[mid+1]) {
			return nums[mid]
		}
		if ((mid+1)%2 == 0 && nums[mid] != nums[mid-1]) || ((mid+1)%2 != 0 && nums[mid] != nums[mid+1]) {
			end = mid
		} else {
			start = mid

		}
	}
	return nums[start]
}
