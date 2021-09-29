func numWays(steps int, arrLen int) int {
    const mod = 1e9 + 7
    arrLen = min(arrLen, steps / 2 + 1)
    prev := make([]int, arrLen)
    dp := make([]int, arrLen)
    dp[0] = 1
    for step := 1; step <= steps; step++ {
        dp, prev = prev, dp
        maxLoc := min(step + 1, arrLen)
        dp[0] = (prev[0] + prev[1]) % mod
        for loc := 1; loc < maxLoc - 1; loc++ {
            dp[loc] = (prev[loc - 1] + prev[loc] + prev[loc + 1]) % mod
        }
        dp[maxLoc - 1] = (prev[maxLoc - 2] + prev[maxLoc - 1]) % mod
    }
    return dp[0]
}

func min(lhs int, rhs int) int {
    if rhs < lhs { return rhs }
    return lhs
}
