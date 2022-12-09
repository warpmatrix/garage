func max(lhs int, rhs int) int {
    if lhs > rhs { return lhs }
    return rhs
}

func _rob(nums [] int) int {
    pre, cur := 0, 0
    for _, v := range nums {
        pre, cur = cur, max(cur, pre + v)
    }
    return cur
}

func rob(nums []int) int {
    if len(nums) == 1 { return nums[0] }
    return max(_rob(nums[:len(nums) - 1]), _rob(nums[1:]))
}