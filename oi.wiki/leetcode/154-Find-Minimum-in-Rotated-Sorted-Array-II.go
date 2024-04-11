func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }
    return rhs
}

func findMin(nums []int) int {
	lptr, rptr := 0, len(nums) - 1
	for lptr < rptr {
		mid := (lptr + rptr) / 2
		if nums[mid] > nums[rptr] {
			lptr = mid + 1
		} else if nums[mid] < nums[rptr] {
			rptr = mid
		} else {
			rptr--
		}
	}
	return nums[rptr]
}

// func findMin(nums []int) int {
//     min := nums[0]
//     for i := len(nums) - 1; i >= 0; i-- {
//         if nums[i] > min {
//             break
//         }
//         min = nums[i]
//     }
//     return min
// }
