func totalHammingDistance(nums []int) int {
    // m := max(nums)
    ret := 0
    // for i := 0; (m >> i) > 0; i++ {
    for i := 0; i < 30; i++ {
        cnt1 := 0
        for _, num := range nums {
            cnt1 += (num >> i) & 1
        }
        ret += (len(nums) - cnt1) * cnt1
    }
    return ret
}

// func max(nums []int) int {
//     ret := math.MinInt32
//     for _, num := range nums {
//         if ret < num { ret = num }
//     }
//     return ret
// }
