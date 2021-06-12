func largestNumber(costs []int, target int) string {
    digLens := make([]int, target + 1)
    for i := range digLens {
        digLens[i] = math.MinInt32
    }
    digLens[0] = 0
    for _, cost := range costs {
        for subCost := cost; subCost <= target; subCost++ {
            digLens[subCost] = max(digLens[subCost], digLens[subCost - cost] + 1)
        }
    }
    if digLens[target] < 0 { return "0" }
    ret := make([]byte, 0, digLens[target])
    subCost := target
    for i := len(costs) - 1; i >= 0; i-- {
        for subCost >= costs[i] && 
            digLens[subCost] == digLens[subCost - costs[i]] + 1 {
            subCost -= costs[i]
            ret = append(ret, byte('1' + i))
        }
    }
    return string(ret)
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
    // dp := make([][]string, len(costs))
    // for i := range dp {
    //     dp[i] = make([]string, target + 1)
    // }
    // for i := range costs {
    //     cost = costs[len(costs) - i - 1]
    //     for subCost := cost; subCost <= target; subCost++ {
    //         for cnt := 1; cnt * cost < subCost; cnt++ {
    //             dp[i][subCost] = max(dp[i][subCost], dp[i - 1][subCost - cost])
    //         }
    //     }
    // }
    // return ""
// [cost] string
// digMax = tar/min

// func max(lhs, rhs string) string {
//     if len(lhs) > len(rhs) { return lhs }
//     if len(lhs) < len(rhs) { return rhs }
//     if rhs > lhs { return rhs }
//     return lhs
// }
