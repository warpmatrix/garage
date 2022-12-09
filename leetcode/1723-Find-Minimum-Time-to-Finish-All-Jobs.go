func minimumTimeRequired(times []int, k int) int {
    const inf = math.MaxInt32
    setNum := 1 << len(times)
    sumTimes := make([]int, setNum)
    for i := 1; i < len(sumTimes); i++ {
        x := bits.TrailingZeros(uint(i))
        y := i ^ 1 << x
        sumTimes[i] = sumTimes[y] + times[x]
    }
    minTimes := make([][]int, k)
    for i := range minTimes {
        minTimes[i] = make([]int, setNum)
    }
    for subSet, time := range sumTimes {
        minTimes[0][subSet] = time
    }
    for wNum := 1; wNum < k; wNum++ {
        for set := 0; set < setNum; set++ {
            minTime := inf
            for subSet := set; subSet > 0; subSet = (subSet-1) & set {
                minTime = min(minTime, max(minTimes[wNum-1][set-subSet], sumTimes[subSet]))
            }
            minTimes[wNum][set] = minTime
        }
    }
    return minTimes[k - 1][setNum - 1]
}

func min(lhs int, rhs int) int {
    if rhs < lhs { return rhs }
    return lhs
}

func max(lhs int, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
