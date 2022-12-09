func strangePrinter(s string) int {
    dp := make([][]int, len(s))
    for i := range dp {
        dp[i] = make([]int, len(s))
        dp[i][i] = 1
    }
    for l := 1; l < len(s); l++ {
        for head := 0; head + l < len(s); head++  {
            if s[head] == s[head + l] {
                dp[head][head + l] = dp[head][head + l - 1]
                continue
            }
            dp[head][head + l] = math.MaxInt32
            for delim := 0; delim < l; delim++ {
                dp[head][head + l] = min(dp[head][head + l],
                    dp[head][head + delim] + dp[head + delim + 1][head + l])
            }
        }
    }
    return dp[0][len(s) - 1]
}

func min(lhs, rhs int) int {
    if rhs < lhs { return rhs }
    return lhs
}
