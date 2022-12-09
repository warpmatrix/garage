func titleToNumber(columnTitle string) int {
    ret := 0
    for _, ch := range []byte(columnTitle) {
        ret = ret * 26 + int(ch - 'A') + 1
    }
    return ret
}
