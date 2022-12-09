func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    cnts := map[int]int{}
    for _, num := range arr {
        cnts[min(len(arr), num)]++
    }
    miss := 0
    for i := 1; i <= len(arr); i++ {
        if cnts[i] == 0 {
            miss++
        } else {
            miss -= min(cnts[i] - 1, miss)
        }
    }
    return len(arr) - miss
}

func min(lhs, rhs int) int {
    if rhs < lhs { return rhs }
    return lhs
}
