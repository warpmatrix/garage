func numDecodings(s string) int {
    if len(s) == 0 { return 1 }
    if s[0] == '0' { return 0 }
    p2, p1, ret := 0, 0, 1
    for i := 0; i < len(s); i++ {
        p2, p1 = p1, ret
        if s[i] == '0' { ret = 0 }
        if i > 0 {
            num, _ := strconv.Atoi(s[i-1:i+1])
            if num >= 10 && num <= 26 {
                ret += p2
            }
        }
    }
    return ret
}
