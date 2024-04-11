func minPairSum(nums []int) int {
    sort.Ints(nums)
    ret := 0
    for i := range nums[:len(nums) / 2] {
        ret = max(ret, nums[i] + nums[len(nums) - i - 1])
    }
    return ret
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
