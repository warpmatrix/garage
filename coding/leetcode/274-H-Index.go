func hIndex(nums []int) int {
    cnts := make([]int, len(nums) + 1)
    for _, num := range nums {
        if num >= len(nums) {
            cnts[len(nums)]++
        } else {
            cnts[num]++
        }
    }
    sum, ret := 0, 0
    for i := len(nums); i >= 0; i-- {
        sum += cnts[i]
        if sum >= i {
            ret = i
            break
        }
    }
    return ret
}
