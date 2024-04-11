func getNext(pattern string) []int {
    j, k := 0, -1
    next := make([]int, len(pattern))
    next[0] = -1
    for j < len(pattern) - 1 {
        if k == -1 || pattern[j] == pattern[k] {
            j, k = j + 1, k + 1
            next[j] = next[k]
        } else {
            k = next[k]
        }
    }
    return next
}

func strStr(haystack string, needle string) int {
    if needle == "" { return 0 }
    next := getNext(needle)
    strIdx, pIdx := 0, 0
    for strIdx < len(haystack) && pIdx < len(needle) {
        if pIdx == -1 || haystack[strIdx] == needle[pIdx] {
            strIdx, pIdx = strIdx + 1, pIdx + 1
        } else {
            pIdx = next[pIdx]
        }
    }
    if pIdx == len(needle) { return strIdx - pIdx }
    return -1
}
