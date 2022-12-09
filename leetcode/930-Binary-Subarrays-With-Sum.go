func numSubarraysWithSum(nums []int, goal int) int {
    oneLocs := []int{0}
    for i, num := range nums {
        if num == 1 {
            oneLocs = append(oneLocs, i + 1)
        }
    }
    oneLocs = append(oneLocs, len(nums) + 1)
    ret := 0
    for head := 1; head < len(oneLocs) - goal; head++ {
        if goal == 0 {
            tmp := oneLocs[head] - oneLocs[head - 1]
            ret += tmp * (tmp - 1) / 2
        } else {
            tail := head + goal - 1
            ret += (oneLocs[head] - oneLocs[head - 1]) * (oneLocs[tail + 1] - oneLocs[tail])
        }
    }
    return ret
}
