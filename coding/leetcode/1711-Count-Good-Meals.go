func countPairs(nums []int) int {
    const mod = 1e9 + 7
    maxNum, cnts := 0, map[int]int{}
    for _, num := range nums {
        maxNum = max(maxNum, num)
        cnts[num]++
    }

    ret, l:= 0, bits.Len(uint(maxNum))
    calced := map[int]int{}
    for num, cnt := range cnts {
        for i := 0; i <= l; i++ {
            sum := 1 << i
            ret += cnt * calced[sum - num] % mod
            ret %= mod
        }
        calced[num] = cnt
        if isPower2(num) {
            ret += (cnt * (cnt - 1) / 2) % mod
            ret %= mod
        }
    }
    return ret
}

func isPower2(n int) bool{
    return n > 0 && (n &(n-1)) == 0
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
