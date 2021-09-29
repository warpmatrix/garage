func decode(encodeds []int) []int {
    odd := 0
    for i := 1; i < len(encodeds); i += 2 {
        odd ^= encodeds[i]
    }
    ret := make([]int, len(encodeds) + 1)
    ret[0] = odd ^ xor(len(ret))
    for i, v := range encodeds {
        ret[i + 1] = ret[i] ^ v
    }
    return ret
}

func xor(num int) int {
    switch num % 4 {
    case 0:
        return num
    case 1:
        return 1
    case 2:
        return num + 1
    // case 3:
    default:
        return 0
    }
}

// func decode(encoded []int) []int {
//     tmps := make([]int, len(encoded) + 1)
//     tmp := 0
//     for i := range tmps {
//         if i == 0 { continue }
//         tmps[i] = tmps[i - 1] ^ encoded[i - 1]
//         tmp = tmp ^ tmps[i]
//     }
//     tmps[0] = tmp ^ xor(len(tmps))
//     for i := range tmps {
//         if i == 0 { continue }
//         tmps[i] ^= tmps[0]
//     }
//     return tmps
// }
