func minCost(houses []int, cost [][]int, m int, n int, target int) int {
    const inf = math.MaxInt32 << 1
    dp := make([][][]int, m)
    // house index
    for hIdx := range dp {
        dp[hIdx] = make([][]int, n)
        for color := range dp[hIdx] {
            dp[hIdx][color] = make([]int, target)
            for neiNum := range dp[hIdx][color] {
                dp[hIdx][color][neiNum] = inf
            }
        }
        houses[hIdx]--
    }
    // house index
    for hIdx := range dp {
        if hIdx == 0 {
            for color := range dp[hIdx] {
                if houses[hIdx] != -1 && houses[hIdx] != color { continue }
                if houses[hIdx] == -1 {
                    dp[hIdx][color][0] = cost[hIdx][color]
                // house[hIdx] == color
                } else {
                    dp[hIdx][color][0] = 0
                }
            }
            continue
        }
        for color := range dp[hIdx] {
            if houses[hIdx] != -1 && houses[hIdx] != color { continue }
            // neiborhood number
            for neiNum := range dp[hIdx][color] {
                for preCol := range dp[hIdx] {
                    if color == preCol {
                        dp[hIdx][color][neiNum] = min(dp[hIdx][color][neiNum],
                            dp[hIdx - 1][color][neiNum])
                    } else if neiNum > 0 {
                        dp[hIdx][color][neiNum] = min(dp[hIdx][color][neiNum],
                            dp[hIdx - 1][preCol][neiNum - 1])
                    }
                }
                if houses[hIdx] == -1 && dp[hIdx][color][neiNum] != inf {
                    dp[hIdx][color][neiNum] += cost[hIdx][color]
                }
            }
        }
    }
    ret := inf
    for _, v := range dp[m - 1] {
        ret = min(ret, v[target - 1])
    }
    if ret == inf { ret = -1 }
    return ret
}

func min(lhs int, rhs int) int {
    if rhs < lhs { return rhs }
    return lhs
}
