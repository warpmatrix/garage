func frequencySort(s string) string {
    cnts := make(map[byte]int)
    for _, ch := range s {
        cnts[byte(ch)]++
    }
    maxCnt := 0
    for _, cnt := range cnts {
        maxCnt = max(maxCnt, cnt)
    }
    buckets := make([][]byte, maxCnt)
    for ch, cnt := range cnts {
        buckets[cnt - 1] = append(buckets[cnt - 1], ch)
    }

    ret := []byte{}
    for i := len(buckets) - 1; i >= 0; i-- {
        for _, ch := range buckets[i] {
            ret = append(ret, bytes.Repeat([]byte{ch}, i + 1)...)
        }
    }
    return string(ret)
}

func max(lhs, rhs int) int {
    if rhs > lhs { return rhs }
    return lhs
}
