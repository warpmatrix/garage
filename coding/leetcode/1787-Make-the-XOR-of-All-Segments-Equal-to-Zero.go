func minChanges(nums []int, k int) int {
    const max, inf = 1 << 10, 0x3f3f3f3f
    prev, dp := []int{}, make([]int, max)
    dp[0] = 0
    for i := range dp[1:] {
        dp[1:][i] = inf
    }
    prevMinDp, minDp := 0, 0
    for i := 0; i < k; i++ {
        prev, prevMinDp, minDp = append([]int{}, dp...), minDp, inf
        cnts := map[int]int{}
        size := (len(nums) + k - i - 1) / k
        for j := i; j < len(nums); j += k {
            cnts[nums[j]]++
        }
        for mask := range dp {
            dp[mask] = prevMinDp
            for num, cnt := range cnts {
                dp[mask] = min(prev[mask ^ num] - cnt, dp[mask])
            }
            dp[mask] += size
            minDp = min(minDp, dp[mask])
        }
    }
    return dp[0]
}

func min(lhs int, rhs int) int {
    if rhs < lhs { return rhs }
    return lhs
}
