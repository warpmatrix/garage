func min(lhs int, mid int, rhs int) int {
    if lhs <= mid && lhs <= rhs {
        return lhs
    }
    if mid <= rhs && mid <= rhs {
        return mid
    }
    return rhs
}

func nthUglyNumber(n int) int {
    p2, p3, p5 := 0, 0, 0
    uglyNums := make([]int, n)
    uglyNums[0] = 1
    for i := 1; i < n; i++ {
        uglyNums[i] = min(uglyNums[p2]*2, uglyNums[p3]*3, uglyNums[p5]*5)
        if uglyNums[i] == uglyNums[p2] * 2 {
            p2++
        }
        if uglyNums[i] == uglyNums[p3] * 3 {
            p3++
        }
        if uglyNums[i] == uglyNums[p5] * 5 {
            p5++
        }
    }
    return uglyNums[n - 1]
}