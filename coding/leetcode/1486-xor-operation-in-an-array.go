func sumXor(x int) int {
    switch x % 4 {
    case 0:
        return x
    case 1:
        return 1
    case 2:
        return x + 1
	// case 3:
    default:
        return 0
    }
}

func xorOperation(n, start int) (ans int) {
    s, e := start>>1, n&start&1
    ret := sumXor(s-1) ^ sumXor(s+n-1)
    return ret<<1 | e
}

// func xorOperation(n int, start int) int {
//     ret := 0
//     for i := 0; i < n; i++ {
//         ret ^= start + 2*i
//     }
//     return ret
// }
