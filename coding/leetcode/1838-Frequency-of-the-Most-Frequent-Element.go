func maxFrequency(nums []int, k int) int {
    ret := 1
    sort.Ints(nums)
    for lIdx, rIdx, sum := 0, 1, 0; rIdx < len(nums); rIdx++ {
        sum += (nums[rIdx] - nums[rIdx - 1]) * (rIdx - lIdx)
        for sum > k {
            sum -= nums[rIdx] - nums[lIdx]
            lIdx++
        }
        ret = max(ret, rIdx - lIdx + 1)
    }
    return ret
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
