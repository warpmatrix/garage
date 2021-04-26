func calcDay(weights []int, cap int) int {
    day, sum := 1, 0
    for _, weight := range weights {
        sum = sum + weight
        if sum > cap {
            sum, day = weight, day + 1
        }
    }
    return day
}

func shipWithinDays(weights []int, D int) int {
    minCap, maxCap := 0, 0
    for _, weight := range weights {
        if minCap < weight { minCap = weight }
        maxCap += weight
    }
    for minCap < maxCap {
        cap := (minCap + maxCap) / 2
        day := calcDay(weights, cap)
        if day > D {
            minCap = cap + 1
        } else {
            maxCap = cap 
        }
    }
    return maxCap
}
