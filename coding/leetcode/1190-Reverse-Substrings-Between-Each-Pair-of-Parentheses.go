func reverseParentheses(s string) string {
    matchIdx := make([]int, len(s))
    st := []int{}
    for i, ch := range s {
        if ch == '(' {
            st = append(st, i)
        } else if ch == ')' {
            lIdx := st[len(st) - 1]
            st = st[:len(st) - 1]
            matchIdx[lIdx], matchIdx[i] = i, lIdx
        }
    }
    ret := []byte{}
    for i, det := 0, 1; i < len(s); i += det {
        if s[i] == '(' || s[i] == ')' {
            i = matchIdx[i]
            det = -det
        } else {
            ret = append(ret, s[i])
        }
    }
    return string(ret)
}
