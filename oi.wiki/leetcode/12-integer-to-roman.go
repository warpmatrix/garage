func intToRoman(num int) (ret string) {
    syms := []string{"I", "V", "X", "L", "C", "D", "M"}
    m := make([][]string, (len(syms) + 1) / 2)
    digIdx, idx := 0, 0
    for _, sym := range syms {
        if idx == 0 { m[digIdx] = make([]string, 2) }
        m[digIdx][idx] = sym
        idx++
		digIdx, idx = digIdx + idx / 2, idx % 2
    }
    parseNum := func(num int, digIdx int) (ret string) {
        if num == 0 { return }
        cnt := num
        if num >= 4 && num <= 8 {
            ret = m[digIdx][1]
            cnt = num - 5
        } else if num > 8 {
            ret = m[digIdx + 1][0]
            cnt = num - 10
        }
        if cnt == -1 {
            return m[digIdx][0] + ret
        }
        for i := 0; i < cnt; i++ {
            ret = ret + m[digIdx][0]
        }
        return
    }
    digIdx = 0
    for num != 0 {
        ret = parseNum(num % 10, digIdx) + ret
        num, digIdx = num / 10, digIdx + 1
    }
    return
}
