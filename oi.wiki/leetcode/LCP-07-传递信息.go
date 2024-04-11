func numWays(n int, relation [][]int, k int) int {
    wayNums := make([]int, n)
    wayNums[0] = 1
    for i := 0; i < k; i++ {
        tmp := make([]int, n)
        for _, edge := range relation {
            src, dest := edge[0], edge[1]
            tmp[dest] += wayNums[src]
        }
        wayNums = tmp
    }
    return wayNums[n - 1]
}
