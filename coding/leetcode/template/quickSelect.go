func findKth(nums []int, k int) int {
    if k + 1 > len(nums) { return -1 }
    if len(nums) == 1 { return nums[0] }
    if len(nums) == 2 {
        if k == 0 { return min(nums[0], nums[1]) }
        return max(nums[0], nums[1])
    }
    pivot := getPivot(nums)
    lIdx, rIdx := 0, len(nums) - 2
    for lIdx < rIdx {
        for lIdx < rIdx && nums[lIdx] < pivot {
            lIdx++
        } 
        for lIdx < rIdx && nums[rIdx] >= pivot {
            rIdx--
        }
        if lIdx < rIdx {
            nums[rIdx], nums[lIdx] = nums[lIdx], nums[rIdx]
        }
    }
    // jump out of loop, the nums[lIdx] doesn't jedge yet, there are two cases:
    // 1. nums[lIdx] < pivot  2. nums[lIdx] >= pivot
    if nums[lIdx] < pivot { lIdx++ }
    nums[lIdx], nums[len(nums) - 1] = nums[len(nums) - 1], nums[lIdx]
    if lIdx < k { return findKth(nums[lIdx+1:], k-(lIdx+1)) }
    if lIdx > k { return findKth(nums[:lIdx], k) }
    return pivot
}

func getPivot(nums []int) int {
    mid := len(nums) / 2
    if (nums[len(nums) - 1] - nums[0]) * (nums[mid] - nums[0]) <= 0 {
        nums[0], nums[len(nums) - 1] = nums[len(nums) - 1], nums[0]
        // return nums[0]
    }
    if (nums[len(nums) - 1] - nums[mid]) * (nums[0] - nums[mid]) <= 0 {
        nums[mid], nums[len(nums) - 1] = nums[len(nums) - 1], nums[mid]
        // return nums[mid]
    }
    return nums[len(nums) - 1]
}

func min(lhs int, rhs int) int {
    if rhs < lhs { return rhs }
    return lhs
}

func max(lhs int, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
