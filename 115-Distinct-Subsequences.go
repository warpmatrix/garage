func numDistinct(s string, t string) int {
    if len(s) < len(t) {
        return 0;
    }
    copy := []rune(s)
    for i := 0; i < len(copy); i++ {
        if copy[i] >= 'a' && copy[i] <= 'z' {
            copy[i] = copy[i] - 'a' + 'A' + 26
        }
    }
    s = string(copy)
    copy = []rune(t)
    for i := 0; i < len(copy); i++ {
        if copy[i] >= 'a' && copy[i] <= 'z' {
            copy[i] = copy[i] - 'a' + 'A' + 26
        }
    }
    t = string(copy)
    locs := make([][]int, 52)
    for i := 0; i < len(s); i++ {
        locs[s[i] - 'A'] = append(locs[s[i] - 'A'], i)
    }
    dp := make([][]int, 52)
    for i := len(t) - 1; i >= 0; i-- {
        dp[i] = make([]int, len(locs[t[i] - 'A']))
        for j := 0; j < len(locs[t[i] - 'A']); j++ {
            if i == len(t) - 1 {
                dp[i][j] = 1
                continue
            }
            for k := len(locs[t[i + 1] - 'A']) - 1; k >= 0; k-- {
                if locs[t[i] - 'A'][j] >= locs[t[i + 1] - 'A'][k] {
                   break 
                }
                dp[i][j] += dp[i + 1][k]
            }
        }
    }
    ans := 0
    for i := 0; i < len(dp[0]); i++ {
        ans += dp[0][i]
    }
    return ans
}

// func numDistinct(s string, t string) int {
//     if len(s) < len(t) {
//         return 0
//     }
//     dp := make([][]int, len(s) + 1)
//     for i := 0; i < len(dp); i++ {
//         dp[i] = make([]int, len(t) + 1)
//     }
//     for i := 0; i <= len(s); i++ {
//         dp[i][len(t)] = 1;
//     }
//     for i := len(s) - 1; i >= 0; i-- {
//         for j := len(t) - 1; j >= 0; j-- {
//             if (s[i] == t[j]) {
//                 dp[i][j] = dp[i + 1][j + 1] + dp[i + 1][j];
//             } else {
//                 dp[i][j] = dp[i + 1][j];
//             }
//         }
//     }
//     return dp[0][0];
// }