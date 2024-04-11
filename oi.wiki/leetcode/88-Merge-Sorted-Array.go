func merge(nums1 []int, m int, nums2 []int, n int)  {
	ptr1, ptr2, idx := m - 1, n - 1, m + n - 1
	for ptr1 >= 0 || ptr2 >= 0 {
		if ptr1 >= 0 && ptr2 >= 0 {
			if nums2[ptr2] > nums1[ptr1] {
				nums1[idx] = nums2[ptr2]
				ptr2--
			} else {
				nums1[idx] = nums1[ptr1]
				ptr1--
			}
			idx--
		} else {
			for ptr1 >= 0 {
				nums1[idx] = nums1[ptr1]
				ptr1--
			}
			for ptr2 >= 0 {
				nums1[idx] = nums2[ptr2]
				ptr2--
			}
		}
	}
}