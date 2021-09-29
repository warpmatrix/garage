func subarraySum(nums []int, k int) int {
    cnts := map[int]int{0:1}
    preSum, ret := 0, 0
    for _, num := range nums {
        preSum += num
        ret += cnts[preSum - k]
        cnts[preSum]++
    }
    return ret
}
