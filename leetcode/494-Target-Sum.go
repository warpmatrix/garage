func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if !(sum >= target && (sum - target) & 1 == 0) {
        return 0
    }

    negSum := (sum - target) >> 1
    dp := make([]int, negSum + 1)
    dp[0] = 1
    for _, num := range nums {
        for sum := negSum; sum >= num; sum-- {
            dp[sum] += dp[sum - num]
        }
    }
    return dp[negSum]
}

// func findTargetSumWays(nums []int, target int) int {
//     var cnts, pCnts map[int]int
//     cnts = map[int]int { 0:1 }
//     for _, num := range nums {
//         pCnts, cnts = cnts, make(map[int]int)
//         for sum, cnt := range pCnts {
//             cnts[sum + num] += cnt
//             cnts[sum - num] += cnt
//         }
//     }
//     return cnts[target]
// }
