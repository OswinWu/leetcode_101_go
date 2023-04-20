package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Index := m - 1
	nums2Index := n - 1
	resultIndex := m + n - 1
	if resultIndex > 0 {
		for nums1Index >= 0 && nums2Index >= 0 {
			if nums1[nums1Index] < nums2[nums2Index] {
				nums1[resultIndex] = nums2[nums2Index]
				nums2Index--
			} else {
				nums1[resultIndex] = nums1[nums1Index]
				nums1Index--
			}
			resultIndex--
		}
	} else {
		if n == 0 {
			return
		} else {
			nums1[0] = nums2[0]
		}
	}
	for nums2Index >= 0 {
		nums1[resultIndex] = nums2[nums2Index]
		resultIndex--
		nums2Index--
	}
}
