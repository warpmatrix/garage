func maxPoints(points [][]int) int {
    maxCnt := 0
    for i, p1 := range points {
        if maxCnt * 2 >= len(points) || maxCnt >= len(points) - i {
            break
        }
        // map detX and detY into the count of slope
        slopeCnts := map[int]map[int]int{}
        for _, p2 := range points[i+1:] {
            divisor, dividend := p2[0] - p1[0], p2[1] - p1[1]
            dividend, divisor = simplFrac(dividend, divisor)
            if _, ok := slopeCnts[dividend]; !ok {
                slopeCnts[dividend] = map[int]int{}
            }
            slopeCnts[dividend][divisor]++
            maxCnt = max(maxCnt, slopeCnts[dividend][divisor])
        }
    }
    return maxCnt + 1
}

func simplFrac(dividend, divisor int) (int, int) {
    if divisor == 0 { return 1, 0 }
    if dividend == 0 { return 0, 1 }
    if divisor < 0 {
        dividend, divisor = -dividend, -divisor
    }
    gcd := calcGcd(abs(dividend), abs(divisor))
    dividend /= gcd
    divisor /= gcd
    return dividend, divisor
}

func calcGcd(lhs, rhs int) int {
    for lhs % rhs != 0 {
        lhs, rhs = rhs, lhs % rhs
    }
    return rhs
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}

func abs(num int) int {
    if num < 0 { return -num }
    return num
}
