func combinationSum4(nums []int, target int) int {
    dp := make([]int, target + 1)
    dp[0] = 1
    for i := 1; i <= target; i++ {
        for _, v := range nums {
            if i >= v {
                dp[i] += dp[i - v]
            }
        }
    }
    return dp[target]
}

// func combinationSum4(nums []int, target int) int {
//     comb := map[int]int{0:1}
//     var memSrch func(tgt int) int
//     memSrch = func(tgt int) int {
//         if tgt < 0 { return 0 }
//         if v, ok := comb[tgt]; ok { return v }
//         sum := 0
//         for _, v := range nums {
//             sum = sum + memSrch(tgt - v)
//         }
//         comb[tgt] = sum
//         return sum
//     }
//     return memSrch(target)
// }