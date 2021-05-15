func romanToInt(s string) int {
    m := map[byte]int{
        'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000
    }
    ret := 0
    for i := range s {
        if i < len(s) - 1 && m[s[i]] < m[s[i + 1]] {
            ret -= m[s[i]]
        } else {
            ret += m[s[i]]
        }
    } 
    return ret
}
