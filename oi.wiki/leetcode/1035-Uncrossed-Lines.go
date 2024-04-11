func maxUncrossedLines(nums1 []int, nums2 []int) int {
    dp := make([][]int, len(nums1) + 1)
    for i := range dp {
        dp[i] = make([]int, len(nums2) + 1)
    }
    dp[0][0] = 0
    for i1, num1 := range nums1 {
        for i2, num2 := range nums2 {
            if num1 == num2 {
                dp[i1 + 1][i2 + 1] = dp[i1][i2] + 1
            } else {
                dp[i1 + 1][i2 + 1] = max(dp[i1][i2 + 1], dp[i1 + 1][i2])
            }
        }
    }
    return dp[len(nums1)][len(nums2)]
}

func max(lhs int, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
