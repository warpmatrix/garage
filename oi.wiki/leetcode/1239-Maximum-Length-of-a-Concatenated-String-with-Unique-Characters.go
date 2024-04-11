func maxLength(arr []string) int {
    ret := 0
    stats := map[int]struct{}{
        0 : struct{}{},
    }
    for _, str := range arr {
        num := parseNum(str)
        for stat := range stats {
            if num & stat == 0 {
                stats[num | stat] = struct{}{}
                ret = max(ret, bits.OnesCount(uint(num | stat)))
            }
        }
    }
    return ret
}

func parseNum(str string) int {
    ret := 0
    for _, c := range str {
        nBit := 1 << (c - 'a')
        if ret & nBit != 0 {
            return 0
        }
        ret |= nBit
    }
    return ret
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
