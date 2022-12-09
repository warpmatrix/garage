func reverse(x int) int {
    ret := 0
    for x != 0 {
        if ret < math.MinInt32 / 10 || ret > math.MaxInt32 / 10 {
            return 0
        }
        ret = ret * 10 + x % 10
        x /= 10
    }
    return ret
}
