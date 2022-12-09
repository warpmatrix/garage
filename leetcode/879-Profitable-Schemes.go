func profitableSchemes(n int, minProfit int, groups []int, profits []int) int {
    const mod = 1e9 + 7
    dp := make([][]int, n + 1)
    for memNum := range dp {
        dp[memNum] = make([]int, minProfit + 1)
    }
    dp[0][0] = 1
    for i, memReq := range groups {
        for memNum := n; memNum >= memReq; memNum-- {
            for profit := minProfit; profit >= 0; profit-- {
                dp[memNum][profit] = (dp[memNum][profit] +
                    dp[memNum - memReq][max(0, profit - profits[i])]) % mod
            }
        }
    }
    ret := 0
    for _, cnt := range dp {
        ret = (ret + cnt[minProfit]) % mod
    }
    return ret
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
