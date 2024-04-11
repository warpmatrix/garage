func triangleNumber(nums []int) int {
    sort.Ints(nums)
    ret := 0
    for i := range nums {
        k := 2
        for j := 1; i + j < len(nums); j++ {
            for i + k < len(nums) && nums[i + k] < nums[i] + nums[i + j] {
                k++
            }
            ret += max(k - j - 1, 0)
        }
    }
    return ret
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
