func hammingDistance(x int, y int) int {
    return hammingWeight(uint32(x ^ y))
}

func hammingWeight(num uint32) int {
    ret := 0
    for num != 0 {
        num &= num - 1
        ret++
    }
    return ret
}
