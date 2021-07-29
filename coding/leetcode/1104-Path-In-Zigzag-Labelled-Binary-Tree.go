func pathInZigZagTree(label int) []int {
    conv := func (num int) int {
        if depth := bits.Len(uint(num)); depth & 1 == 1 {
            return num
        } else {
            return (1 << depth) - num + (1 << (depth - 1)) - 1
        }
    }
    loc := conv(label)
    ret := make([]int, bits.Len(uint(label)))
    for i := len(ret) - 1; i >= 0; i-- {
        ret[i] = conv(loc)
        loc = loc >> 1
    }
    return ret
}
