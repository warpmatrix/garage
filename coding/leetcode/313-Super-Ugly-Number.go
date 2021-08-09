func nthSuperUglyNumber(n int, primes []int) int {
    ptrs := make([]int, len(primes))
    uglyNums := make([]int, n)
    uglyNums[0] = 1
    for i := 1; i < n; i++ {
        uglyNums[i] = math.MaxInt32
        for idx, ptr := range ptrs {
            uglyNums[i] = min(uglyNums[i], uglyNums[ptr] * primes[idx])
        }
        for idx, ptr := range ptrs {
            if uglyNums[i] == uglyNums[ptr] * primes[idx] {
                ptrs[idx]++
            }
        }
    }
    return uglyNums[n - 1]
}

func min(lhs, rhs int) int {
    if rhs < lhs { return rhs }
    return lhs
}
