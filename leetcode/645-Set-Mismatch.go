func findErrorNums(nums []int) []int {
    xor := 0
    for i, num := range nums {
        xor ^= (i + 1) ^ num
    }
    lb := xor & (-xor)
    ret := make([]int, 2)
    for i, num := range nums {
        ret[getBit(num, lb)] ^= num
        ret[getBit(i + 1, lb)] ^= i + 1
    }
    for _, num := range nums {
        if num == ret[1] {
            ret = []int{ret[1], ret[0]}
            break
        }
    }
    return ret
}

func getBit(num int, maskBit int) int {
    if  num & maskBit != 0 { return 0 }
    return 1
}
