func stoneGame(piles []int) bool {
    return true
}

// func stoneGame(piles []int) bool {
//     dp := make([]int, len(piles))
//     for i := 0; i < len(piles); i++ {
//         dp[i] = piles[i]
//     }
//     for i := len(piles) - 2; i >= 0; i-- {
//         for j := i + 1; j < len(piles); j++ {
//             dp[j] = max(piles[i] - dp[j], piles[j] - dp[j - 1])
//         }
//     }
//     return dp[len(piles) - 1] > 0
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }
